package main

import (
	"net/http"
	"strconv" // Chuyển đổi chuỗi sang số nguyên (ID)
	"sync"
	"log"
	"os"
	"todo-api-go/models"

	"github.com/gin-contrib/cors" // Import thư viện CORS

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// In-memory "database"
var (
	tasks      = make(map[int]models.Task) // Map để lưu trữ các công việc, key là ID
	nextID     = 1
	tasksMutex sync.Mutex
)

func main() {
	// Tải các biến môi trường từ file .env.
	// Nếu không có file .env, chương trình vẫn chạy bình thường.
	err := godotenv.Load()
	if err != nil {
		log.Println("Không tìm thấy file .env, sử dụng biến môi trường hệ thống.")
	}

	router := gin.Default()

	// Sử dụng biến môi trường cho AllowOrigins
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		// Giá trị mặc định khi chạy local
		frontendURL = "http://localhost:5173"
	}
	// Cấu hình CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{frontendURL} // Cho phép frontend truy cập
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	router.Use(cors.New(config))

	// Định nghĩa các API endpoints
	router.POST("/tasks", createTask)
	router.GET("/tasks", getTasks)
	router.PUT("/tasks/:id", updateStatus)
	router.DELETE("/tasks/:id", deleteTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}

// @Success 201 {object} map[string]int "Trả về ID của công việc vừa tạo"
// @Failure 400 {object} gin.H "Lỗi khi request body không hợp lệ"
// @Router /tasks [post]
func createTask(c *gin.Context) {
	var newTask models.Task
	// Bind JSON request body vào struct newTask
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasksMutex.Lock()         // Khóa mutex để tránh race condition khi ghi vào map
	defer tasksMutex.Unlock() // Đảm bảo mutex được mở khóa khi hàm kết thúc

	newTask.ID = nextID
	newTask.Completed = false
	tasks[newTask.ID] = newTask
	nextID++

	c.JSON(http.StatusCreated, gin.H{"id": newTask.ID})
}

// @Success 200 {array} models.Task "Trả về danh sách công việc"
// @Router /tasks [get]
func getTasks(c *gin.Context) {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	// Chuyển map tasks thành slice để trả về JSON
	var taskList []models.Task
	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	c.JSON(http.StatusOK, taskList)
}

// @Success 200 {object} gin.H "Thông báo cập nhật thành công"
// @Failure 400 {object} gin.H "ID không hợp lệ hoặc request body không hợp lệ"
// @Failure 404 {object} gin.H "Không tìm thấy công việc"
// @Router /tasks/{id} [put]
func updateStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam) // Chuyển đổi ID từ chuỗi sang số nguyên
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updateBody struct {
		Completed bool `json:"completed" binding:"required"`
	}
	if err := c.ShouldBindJSON(&updateBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	task, exists := tasks[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	task.Completed = updateBody.Completed
	tasks[id] = task // Cập nhật lại task trong map

	c.JSON(http.StatusOK, gin.H{"message": "Task status updated successfully"})
}

// @Success 204 "Xóa thành công, không có nội dung trả về"
// @Failure 400 {object} gin.H "ID không hợp lệ"
// @Failure 404 {object} gin.H "Không tìm thấy công việc"
// @Router /tasks/{id} [delete]
func deleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	_, exists := tasks[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	delete(tasks, id)              // Xóa công việc khỏi map
	c.Status(http.StatusNoContent) // HTTP 204 No Content cho yêu cầu xóa thành công
}
