// Package tool 提供了哈希算法的使用示例和测试函数
package tool

import (
	"fmt"
	"log"
)

// ExampleHashPassword 演示如何使用SimpleHash进行密码哈希和验证
// 这是一个完整的示例，展示了哈希密码、存储结果以及验证密码的全过程
func ExampleHashPassword() {
	// 创建使用默认配置的哈希实例
	hasher := NewDefaultHash()

	// 要哈希的密码
	password := "mySecurePassword123"
	
	// 对密码进行哈希处理
	result, err := hasher.HashPassword(password)
	if err != nil {
		log.Fatal("Hashing failed:", err)
	}

	// 打印密码、哈希值和盐值
	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Hash: %s\n", result.Hash)
	fmt.Printf("Salt: %s\n", result.Salt)

	// 验证正确密码
	isValid, err := hasher.VerifyPassword(password, result.Hash, result.Salt)
	if err != nil {
		log.Fatal("Verification failed:", err)
	}

	fmt.Printf("Password is valid: %t\n", isValid)

	// 验证错误密码
	isValid, err = hasher.VerifyPassword("wrongPassword", result.Hash, result.Salt)
	if err != nil {
		log.Fatal("Verification failed:", err)
	}

	fmt.Printf("Wrong password is valid: %t\n", isValid)

	// Output:
	// Password: mySecurePassword123
	// Hash: <hash_value>
	// Salt: <salt_value>
	// Password is valid: true
	// Wrong password is valid: false
}

// RunHashExample 运行哈希算法示例，展示完整的密码哈希和验证流程
// 该函数会打印配置信息、哈希结果和验证结果
func RunHashExample() {
	fmt.Println("=== Hash Algorithm Example ===")
	
	// 创建使用默认配置的哈希实例
	hasher := NewDefaultHash()
	
	// 显示当前使用的哈希配置参数
	config := hasher.GetConfig()
	fmt.Printf("Hash Config:\n")
	fmt.Printf("  Memory: %d KiB\n", config.Memory)
	fmt.Printf("  Iterations: %d\n", config.Iterations)
	fmt.Printf("  Parallelism: %d\n", config.Parallelism)
	fmt.Printf("  Salt Length: %d\n", config.SaltLength)
	fmt.Printf("  Key Length: %d\n", config.KeyLength)
	
	// 要进行哈希处理的示例密码
	password := "examplePassword123"
	fmt.Printf("\nHashing password: %s\n", password)
	
	// 对密码进行哈希处理
	result, err := hasher.HashPassword(password)
	if err != nil {
		log.Fatal("Hashing failed:", err)
	}
	
	// 打印生成的哈希值和盐值
	fmt.Printf("Hash: %s\n", result.Hash)
	fmt.Printf("Salt: %s\n", result.Salt)
	
	// 验证正确密码
	isValid, err := hasher.VerifyPassword(password, result.Hash, result.Salt)
	if err != nil {
		log.Fatal("Verification failed:", err)
	}
	fmt.Printf("Correct password verification: %t\n", isValid)
	
	// 验证错误密码
	isValid, err = hasher.VerifyPassword("wrongPassword", result.Hash, result.Salt)
	if err != nil {
		log.Fatal("Verification failed:", err)
	}
	fmt.Printf("Wrong password verification: %t\n", isValid)
}