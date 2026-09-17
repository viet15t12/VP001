package messages

const (
	Welcome         = "Xin chao, day la V-SSH Manager"
	MenuPrompt      = "Vui long chon mot trong cac chuc nang duoi day:"
	Exit            = "Cre: viet15t12"
	InvalidChoice   = "Lua chon khong hop le, vui long thu lai"
	MesssChoiceMneu = "chon chuc nang: "
)

var MenuItems = []string{
	"[1] Quan ly SSH key\n",
	"[2] Copy SSH key toi host\n",
	"[3] Tao host moi\n",
	"[4] Ket noi SSH\n",
	"[0] Thoat\n",
}

var MenuItemKeyManager = []string{
	"[1] tao keygen\n",
	"[2] xem danh sach key\n",
	"[3] xoa keygen\n",
	"[4] đổi ten keygen\n",
	"[0] thoát\n",
}
