# Penjelasan
Pada kode dibawah untuk struct Area menurut saya sudah sesuai dengan yang diharapkan. Akan tetapi sebaiknya ditambahkan lagi struct baru yaitu AreaRepository yang dimana berguna untuk menampung data inisiasi DB, disini saya menggunakan GORM untuk mengakses database.

### Before
```
type ( 
    Area struct { 
        ID int64 `gorm:"column:id;primaryKey;"`
        AreaValue int64`gorm:"column:area_value"`
        AreaType string`gorm:"column:type"` 
    } 
)
```

### After
```
type (
	AreaRepository struct {
		DB *gorm.DB
	}
	Area struct {
		ID        int64   `gorm:"column:id;primaryKey;"`
		AreaValue float64 `gorm:"column:area_value"`
		AreaType  string  `gorm:"column:type"`
	}
)
```

Pada kode dibawah ada beberapa yang harus diperbaiki:
1. AreaRepository harus diinisiasi dengan *DB*
2. `ar *Model.Area` dihilangkan, tinggal panggil langsung di dalam function tanap perlu jadi param
3. `Var area Area` dihilangkan, karena tidak dibutuhkan dan ada typo ketika deklarasi harusnya var
4. return err pada tiap case seharusnya dibuat pada akhir switch case
5. pada default case, harusnya return fmt.Println("Tipe area tidak ditemukan")


## Before
```
func (_r *AreaRepository) InsertArea(param1 int32, param2 int64, type []string, ar *Model.Area) (err error) { 
    inst := _r.DB.Model(ar) 
    Var area int area = 0 
    switch type { 
        case ‘persegi panjang’: 
            var area := param1 * param2 
            ar.AreaValue = area 
            ar.AreaType = ‘persegi panjang’ 
            err = _r.DB.create(&ar).Error 
            if err != nil { 
                return err 
            } 
            case ‘persegi’: 
                var area = param1 * param2 
                ar.AreaValue = area ar.AreaType = ‘persegi’ 
                err = _r.DB.create(&ar).Error 
                if err != nil { 
                    return err 
                } 
            case segitiga:
                area = 0.5 * (param1 * param2) 
                ar.AreaValue = area 
                ar.AreaType = ‘segitiga’ 
                err = _r.DB.create(&ar).Error 
                if err != nil { 
                    return err 
                } 
            default: 
                ar.AreaValue = 0 
                ar.AreaType = ‘undefined data’ 
                err = _r.DB.create(&ar).Error 
                if err != nil { 
                    return err 
                } 
    } 
}
```
## After
```
func (_r *AreaRepository) InsertArea(param1 int64, param2 int64, typeArea string) (err error) {
	area := new(Area)
	switch typeArea {
		case "segitiga":
			formula := float64(0.5) * float64((param1 * param2))
			area.AreaValue = formula
			area.AreaType = typeArea
		case "persegi panjang", "persegi":
			formula := param1 * param2
			area.AreaValue = float64(formula)
			area.AreaType = typeArea
		default:
			area.AreaValue = 0
			area.AreaType = "undefined data"
	}

	err = _r.DB.Create(&area).Error
	if err != nil {
		return err
	}
	return nil
}
```

Untuk proses pemanggilan function `InsertArea` saya refactor menjadi menggunakan service yang sudah saya buat. Berikut adalah perbedaannya :
## Before
```
err = _u.repository.InsertArea(10, 10, ‘persegi’) 
if err != nil { 
    log.Error().Msg(err.Error()) err = errors.New(en.ERROR_DATABASE) 
    return err 
}
```
## After
```
type service struct {
	repository AreaRepository
}

func (_u service) Service() error {
	err := _u.repository.InsertArea(1, 2, "persegi panjang")
	if err != nil {
		return err
	}
	return nil
}
```

Untuk kode lengkapnya, silahkan buka file `main.go` [di sini](https://github.com/bangadam/backend-test-majoo-area-train/blob/master/main.go)