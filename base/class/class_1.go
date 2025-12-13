package main

import (
	"fmt"
	"time"
)

// 基础用户信息
type BaseUser struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

func (u *BaseUser) GetCreatedDate() string {
	return u.CreatedAt.Format("2006-01-02")
}

func (u *BaseUser) DisplayBasicInfo() {
	fmt.Printf("用户ID: %d, 姓名: %s, 邮箱: %s\n", u.ID, u.Name, u.Email)
}

// 地址信息
type Address struct {
	Province string
	City     string
	District string
	Detail   string
}

func (a Address) GetFullAddress() string {
	return fmt.Sprintf("%s%s%s%s", a.Province, a.City, a.District, a.Detail)
}

// 普通用户
type NormalUser struct {
	BaseUser            // 嵌入基础用户
	Addresses []Address // 多个地址
	Level     int       // 用户等级
}

func (nu *NormalUser) AddAddress(addr Address) {
	nu.Addresses = append(nu.Addresses, addr)
}

// VIP用户
type VIPUser struct {
	BaseUser             // 嵌入基础用户
	Addresses  []Address // 多个地址
	VIPLevel   int       // VIP等级
	Discount   float64   // 专属折扣
	ExpireTime time.Time // VIP到期时间
}

func (vu *VIPUser) IsVIPValid() bool {
	return time.Now().Before(vu.ExpireTime)
}

func (vu *VIPUser) GetDiscount() float64 {
	if vu.IsVIPValid() {
		return vu.Discount
	}
	return 1.0 // 无折扣
}

// 用户服务
type UserService struct{}

func (us *UserService) CreateUser(name, email string, userType string) interface{} {
	base := BaseUser{
		ID:        generateID(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	switch userType {
	case "normal":
		return &NormalUser{
			BaseUser: base,
			Level:    1,
		}
	case "vip":
		return &VIPUser{
			BaseUser:   base,
			VIPLevel:   1,
			Discount:   0.9,                         // 9折
			ExpireTime: time.Now().AddDate(1, 0, 0), // 1年后到期
		}
	default:
		return nil
	}
}

func generateID() int {
	return int(time.Now().Unix())
}

func main() {
	service := &UserService{}

	// 创建不同类型的用户
	normalUser := service.CreateUser("张三", "zhangsan@example.com", "normal").(*NormalUser)
	vipUser := service.CreateUser("李四", "lisi@example.com", "vip").(*VIPUser)

	// 为用户添加地址
	addr := Address{
		Province: "广东省",
		City:     "深圳市",
		District: "南山区",
		Detail:   "科技园123号",
	}

	normalUser.AddAddress(addr)

	// 演示多态性 - 所有用户都有基础信息
	users := []interface{}{normalUser, vipUser}

	for _, user := range users {
		fmt.Println("\n=== 用户信息 ===")
		switch u := user.(type) {
		case *NormalUser:
			u.DisplayBasicInfo()
			fmt.Printf("用户等级: %d\n", u.Level)
			fmt.Printf("注册时间: %s\n", u.GetCreatedDate())
		case *VIPUser:
			u.DisplayBasicInfo()
			fmt.Printf("VIP等级: %d, 折扣: %.1f, 是否有效: %t\n",
				u.VIPLevel, u.Discount, u.IsVIPValid())
		}
	}
}
