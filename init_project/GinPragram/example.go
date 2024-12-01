package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// logrus
var logger *logrus.Logger

func init() {
	// 初始化 logrus
	logger = logrus.New()

	// 配置 rotatelogs 实现按天分割日志
	logWriter, err := rotatelogs.New(
		"logs/gin-%Y-%m-%d.log",                   // 日志文件名格式
		rotatelogs.WithLinkName("logs/gin"),       // 最新日志的软链接
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 日志最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		fmt.Printf("Failed to initialize log file: %v\n", err)
		return
	}

	// 设置 logrus 的输出到文件和控制台
	logger.SetOutput(io.MultiWriter(logWriter, gin.DefaultWriter))
	logger.SetFormatter(&logrus.JSONFormatter{}) // 使用 JSON 格式
	logger.SetLevel(logrus.InfoLevel)            // 设置日志级别
}
func main() {
	// 路由初始化
	r := gin.Default()

	// 中间件：日志记录
	r.Use(LoggingMiddleware())

	// 中间件：鉴权验证
	r.Use(AuthMiddleware())

	// 路由分组：风险评估模块
	riskAssessmentGroup := r.Group("/risk-assessment")
	{
		// 风险评估 - 识别单个采购风险
		riskAssessmentGroup.GET("/identify/:id", func(c *gin.Context) {
			//var requestData map[string]interface{}

			purchaseID := c.Param("id")
			logger.WithField("request", purchaseID).Info("Processing request")
			c.JSON(200, gin.H{
				"message":     "Risk identified",
				"purchase_id": purchaseID,
				"risk_level":  "High", // 模拟返回
			})

		})

		// 风险评估 - 批量识别采购风险
		riskAssessmentGroup.POST("/batch-identify", func(c *gin.Context) {
			var purchaseData []map[string]interface{}
			if err := c.BindJSON(&purchaseData); err != nil {
				c.JSON(400, gin.H{"error": "Invalid input data"})
				return
			}
			c.JSON(200, gin.H{
				"message":   "Batch risk identified",
				"risks":     len(purchaseData),
				"risk_info": purchaseData,
			})
		})
	}

	// 路由分组：风险报告管理模块
	riskReportGroup := r.Group("/risk-report")
	{
		// 风险报告 - 获取所有风险报告
		riskReportGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "List of risk reports",
				"reports": []string{"Report1", "Report2", "Report3"}, // 模拟返回
			})
		})

		// 风险报告 - 新建风险报告
		riskReportGroup.POST("/create", func(c *gin.Context) {
			var reportData map[string]interface{}
			if err := c.BindJSON(&reportData); err != nil {
				c.JSON(400, gin.H{"error": "Invalid input data"})
				return
			}
			c.JSON(201, gin.H{
				"message": "Risk report created",
				"data":    reportData,
			})
		})

		// 风险报告 - 删除指定风险报告
		riskReportGroup.DELETE("/:id", func(c *gin.Context) {
			reportID := c.Param("id")
			c.JSON(200, gin.H{
				"message":   "Risk report deleted",
				"report_id": reportID,
			})
		})
	}

	// 启动服务器
	fmt.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}

// LoggingMiddleware 将日志写入文件，并每日生成新文件
func LoggingMiddleware() gin.HandlerFunc {
	// 设置日志文件分割规则
	logFile, err := rotatelogs.New(
		"logs/gin-%Y-%m-%d.log",                   // 日志文件名格式
		rotatelogs.WithLinkName("logs/gin"),       // 最新日志的软链接
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 日志最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		log.Fatalf("Failed to initialize log file: %v", err)
	}

	// 中间件实现
	return func(c *gin.Context) {
		start := time.Now()

		// 处理请求
		c.Next()

		// 记录日志信息
		log.SetOutput(logFile)
		log.Printf("[%s] %s %s %d %s\n",
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			time.Since(start),
		)
	}
}

// AuthMiddleware 鉴权验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 模拟鉴权逻辑
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort() // 阻止请求继续向下传递
			return
		}

		// 如果鉴权通过，继续执行后续的逻辑
		c.Next()
	}
}
