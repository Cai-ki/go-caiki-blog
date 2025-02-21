package handler

import (
	"github.com/Cai-ki/go-caiki-blog/internal/post"
	"github.com/Cai-ki/go-caiki-blog/internal/service"
)

var postService = post.Service
var CommentService = service.CommentService
var TagService = service.TagService
