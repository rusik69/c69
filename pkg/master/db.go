package master

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/govnocloud/pkg/client"
	"github.com/rusik69/govnocloud/pkg/types"
	"github.com/sirupsen/logrus"
)

// CreateDBHandler handles the create DB request.
func CreateDBHandler(c *gin.Context) {
	db := types.DB{}
	err := c.BindJSON(&db)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	logrus.Println("Creating DB", db.Name)
	dbInfoString, err := ETCDGet("/db/" + db.Name)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if dbInfoString != "" {
		logrus.Error("db with this name already exists")
		c.JSON(400, gin.H{"error": "db with this name already exists"})
		return
	}
	image := types.DBTypes[db.Type].Image
	containerFlavor := types.DBTypes[db.Type].ContainerFlavor
	ctr, err := client.CreateContainer(types.MasterEnvInstance.ListenHost, types.MasterEnvInstance.ListenPort,
		db.Name+"-db", image, containerFlavor)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	db.Container = ctr
	dbString, err := json.Marshal(db)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = ETCDPut("/db/"+db.Name, string(dbString))
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, db)
}

// GetDBHandler handles the get DB request.
func GetDBHandler(c *gin.Context) {
	dbName := c.Param("name")
	dbInfoString, err := ETCDGet("/db/" + dbName)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if dbInfoString == "" {
		logrus.Error("db with this name does not exist")
		c.JSON(400, gin.H{"error": "db with this name does not exist"})
		return
	}
	db := types.DB{}
	err = json.Unmarshal([]byte(dbInfoString), &db)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, db)
}

// DeleteDBHandler handles the delete DB request.
func DeleteDBHandler(c *gin.Context) {
	dbName := c.Param("name")
	dbInfoString, err := ETCDGet("/db/" + dbName)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if dbInfoString == "" {
		logrus.Error("db with this name does not exist")
		c.JSON(400, gin.H{"error": "db with this name does not exist"})
		return
	}
	logrus.Println("Deleting DB", dbName)
	err = client.DeleteContainer(types.MasterEnvInstance.ListenHost, types.MasterEnvInstance.ListenPort, dbName+"-db")
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = ETCDDelete("/db/" + dbName)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}

// ListDBsHandler handles the list DBs request.
func ListDBsHandler(c *gin.Context) {
	dbNames, err := ETCDList("/db")
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	dbs := []types.DB{}
	for _, dbName := range dbNames {
		dbInfoString, err := ETCDGet("/db/" + dbName)
		if err != nil {
			logrus.Error(err.Error())
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		db := types.DB{}
		err = json.Unmarshal([]byte(dbInfoString), &db)
		if err != nil {
			logrus.Error(err.Error())
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		dbs = append(dbs, db)
	}
	c.JSON(200, dbs)
}

// StartDBHandler handles the start DB request.
func StartDBHandler(c *gin.Context) {
	dbName := c.Param("name")
	dbInfoString, err := ETCDGet("/db/" + dbName)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if dbInfoString == "" {
		logrus.Error("db with this name does not exist")
		c.JSON(400, gin.H{"error": "db with this name does not exist"})
		return
	}
	logrus.Println("Starting DB", dbName)
	db := types.DB{}
	err = json.Unmarshal([]byte(dbInfoString), &db)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = client.StartContainer(types.MasterEnvInstance.ListenHost, types.MasterEnvInstance.ListenPort, db.Container.Name)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}

// StopDBHandler handles the stop DB request.
func StopDBHandler(c *gin.Context) {
	dbName := c.Param("name")
	dbInfoString, err := ETCDGet("/db/" + dbName)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if dbInfoString == "" {
		logrus.Error("db with this name does not exist")
		c.JSON(400, gin.H{"error": "db with this name does not exist"})
		return
	}
	logrus.Println("Stopping DB", dbName)
	db := types.DB{}
	err = json.Unmarshal([]byte(dbInfoString), &db)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = client.StopContainer(types.MasterEnvInstance.ListenHost, types.MasterEnvInstance.ListenPort, db.Container.Name)
	if err != nil {
		logrus.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}
