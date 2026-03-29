package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

type Telemetry struct {
    UserID    string    `json:"user_id"`
    Latencies []float64 `json:"latencies"`
    Hash      string    `json:"hash"`
}

func main() {
    r := gin.Default()
    r.POST("/ingest", func(c *gin.Context) {
        var data Telemetry
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(400, gin.H{"error": "Invalid Data"})
            return
        }
        // TODO: Validar Hash SHA-256 aqui
        fmt.Printf("Dados recebidos de %s: %v\n", data.UserID, data.Latencies)
        c.JSON(202, gin.H{"status": "Accepted", "compliance": "NR-1 Ready"})
    })
    r.Run(":8080")
}
