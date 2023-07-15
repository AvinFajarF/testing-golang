package controller

import (
	"github.com/AvinFajarF/config"
	"github.com/AvinFajarF/entity"
	pb "github.com/AvinFajarF/proto"
	"github.com/gin-gonic/gin"
)

var client = pb.NewUserServiceClient(config.RpcClient)

func AddUser(ctx *gin.Context) {
	var user entity.UserEntity

	err := ctx.ShouldBind(&user)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	data := &pb.User{
		Id: user.ID,
		Name: user.Name,
        Email: user.Email,
	}

	res, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		User: data,
	})

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Successfull adding data",
		"book":    res.User,
	})

}

func ShowAllUsers(ctx *gin.Context){
	res, err := client.GetUser(ctx, &pb.GetUserRequest{})

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Successfull show all users",
		"book":    res,
	})

}

func EditUser(ctx *gin.Context) {
	var user entity.UserEntity
	
		err := ctx.ShouldBind(&user)
	
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err,
			})
			return
		}
	
		data := &pb.User{
			Id: user.ID,
			Name: user.Name,
			Email: user.Email,
		}
	
		res, err := client.UpdateUser(ctx, &pb.UpdateUserRequest{
			User: data,
		})
	
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err,
			})
			return
		}
	
		ctx.JSON(201, gin.H{
			"message": "Successfull editing data",
			"book":    res.User,
		})
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := client.DeleteUser(ctx, &pb.DeleteUserRequest{
		Id: id,
	})


	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}


	ctx.JSON(201, gin.H{
			"message": "Successfull delete data",
			"book":    res,
		})


}