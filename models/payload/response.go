package payload

type DashboardUser struct {
	Adoption int `json:"adoption" form:"adoption"`
	Donation int `json:"donation" form:"donation"`
}

type DashboardAdmin struct {
	TotalUser         int `json:"total_user" form:"total_user"`
	TotalAdoption     int `json:"total_adoption" form:"total_adoption"`
	TotalDonation     int `json:"total_donation" form:"total_donation"`
	TotalPetAvailable int `json:"total_pet_available" form:"total_pet_available"`
	TotalPetAdopted   int `json:"total_pet_adopted" form:"total_pet_"`
}

type GetProfil struct {
	Name      string `json:"name" form:"name"`
	Username  string `json:"username" form:"username"`
	Email     string `json:"email" form:"email"`
	Handphone string `json:"handphone" form:"handphone"`
	Alamat    string `json:"alamat" form:"alamat"`
}

type GetPet struct {
	ID        uint   `json:"id" form:"id"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Status    string `json:"status" form:"status"`
	Category  string `json:"category" form:"category"`
	OwnerID   uint   `json:"owner_id" form:"owner_id"`
	Owner     string `json:"owner" form:"owner"`
	Handphone string `json:"handphone" form:"handphone"`
	Alamat    string `json:"alamat" form:"alamat"`
}

type GetDonateList struct {
	ID        uint   `json:"id" form:"id"`
	Deskripsi string `json:"deskripsi" form:"deskrips"`
	Status    string `json:"status" form:"status"`
	Category  string `json:"category" form:"category"`
}
