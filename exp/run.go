package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
    // Создаем контекст для отслеживания сигналов завершения работы
    ctx, cancel := context.WithCancel(context.Background())
    
    defer cancel()

    // Создаем канал для получения сигналов
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

    // Обработчик сигналов
    go func() {
        sig := <-sigCh
        fmt.Println("sdsdsdsd")
        fmt.Printf("Received signal %s, shutting down...\n", sig)
        // Добавьте здесь код для отключения от базы данных
        // Например, закрытие соединения с базой данных или выполнение других очисток
       
        // Здесь можно добавить код для отключения от базы данных
        fmt.Println("Disconnecting from the database...")
        // Например, закрываем соединение с базой данных
        // db.Close()
        // Отменяем контекст, чтобы завершить работу горутин, если они еще активны
  
    }()

    // Здесь можно добавить код запуска сервера

    // Пример ожидания завершения работы главной горутины по контексту
    <-ctx.Done()
    fmt.Println("Server shutdown complete.")
}
