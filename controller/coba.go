package controller

import (
	"errors"
	"fmt"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"

	"net/http"

	inimodel "github.com/gryzlegrizz/BE_TugasBesar/model"
	cek "github.com/gryzlegrizz/BE_TugasBesar/module"
	"github.com/barganakukuhraditya/BOILERPLATE_TUBES/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertParfume godoc
// @Summary Insert data parfume.
// @Description Input data parfume.
// @Tags Parfume
// @Accept json
// @Produce json
// @Param request body Parfume true "Payload Body [RAW]"
// @Success 200 {object} Parfume
// @Failure 400
// @Failure 500
// @Router /insert [post]
func InsertParfume(c *fiber.Ctx) error {
	db := config.Parfumemongoconn
	var parfume inimodel.Parfume
	if err := c.BodyParser(&parfume); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := cek.InsertParfume(db, "parfume",
		parfume.Nama_Parfume,
		parfume.Jenis_Parfume,
		parfume.Merk,
		parfume.Deskripsi,
		parfume.Harga,
		parfume.Thn_Peluncuran,
		parfume.Stok,
		parfume.Ukuran)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

// GetPresensiID godoc
// @Summary Get By ID Data Parfume.
// @Description Ambil per ID data parfume.
// @Tags Parfume
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Parfume
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /parfume/{id} [get]
func GetParfumeID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := cek.GetParfumeFromID(objID, config.Parfumemongoconn, "parfume")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

// GetParfume godoc
// @Summary Get All Data Parfume.
// @Description Mengambil semua data parfume.
// @Tags Parfume
// @Accept json
// @Produce json
// @Success 200 {object} Parfume
// @Router /parfume [get]
func GetParfume(c *fiber.Ctx) error {
	ps := cek.GetAllParfume(config.Parfumemongoconn, "parfume")
	return c.JSON(ps)
}

// UpdateData godoc
// @Summary Update data parfume.
// @Description Ubah data parfume.
// @Tags Parfume
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body ReqParfume true "Payload Body [RAW]"
// @Success 200 {object} Parfume
// @Failure 400
// @Failure 500
// @Router /update/{id} [put]
func UpdateParfume(c *fiber.Ctx) error {
	db := config.Parfumemongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var parfume inimodel.Parfume
	if err := c.BodyParser(&parfume); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = cek.UpdateParfume(objectID, db, "parfume",
		parfume.Nama_Parfume,
		parfume.Jenis_Parfume,
		parfume.Merk,
		parfume.Deskripsi,
		parfume.Harga,
		parfume.Thn_Peluncuran,
		parfume.Stok,
		parfume.Ukuran)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

func DeleteParfumeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = cek.DeleteParfumeByID(objID, config.Parfumemongoconn, "parfume")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}
