@echo off
:: 创建目录结构
mkdir cmd
mkdir internal
mkdir internal\user
mkdir internal\post
mkdir internal\comment
mkdir internal\tag
mkdir internal\auth
mkdir pkg
mkdir pkg\config
mkdir pkg\logger
mkdir pkg\storage
mkdir pkg\validator
mkdir models
mkdir routes
mkdir utils

:: 创建文件
type nul > cmd\main.go
type nul > internal\user\handler.go
type nul > internal\user\service.go
type nul > internal\user\repository.go
type nul > internal\post\handler.go
type nul > internal\post\service.go
type nul > internal\post\repository.go
type nul > internal\comment\handler.go
type nul > internal\comment\service.go
type nul > internal\comment\repository.go
type nul > internal\tag\handler.go
type nul > internal\tag\service.go
type nul > internal\tag\repository.go
type nul > internal\auth\middleware.go
type nul > internal\auth\jwt.go
type nul > pkg\config\config.go
type nul > pkg\logger\logger.go
type nul > pkg\storage\storage.go
type nul > pkg\validator\validator.go
type nul > models\user.go
type nul > models\post.go
type nul > models\comment.go
type nul > models\tag.go
type nul > routes\routes.go
type nul > utils\errors.go

echo Finish
pause