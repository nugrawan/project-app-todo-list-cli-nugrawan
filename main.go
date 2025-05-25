package main

import (
	"flag"
	"fmt"
	"os"
	"projectgo/service"
	"text/tabwriter"
)

func main() {
	// menggunakan flag untuk input langsung dari argumen
	add := flag.Bool("add", false, "Tambah tugas baru")
	title := flag.String("title", "", "Judul tugas")
	desc := flag.String("desc", "", "Deskripsi tugas")

	list := flag.Bool("list", false, "Tampilkan semua tugas")
	search := flag.String("search", "", "Cari berdasarkan kata kunci")
	done := flag.Int("done", 0, "Tandai tugas selesai berdasarkan ID")
	delete := flag.Int("delete", 0, "Hapus tugas berdasarkan ID")

	flag.Parse()

	_ = service.Load()

	switch {
	case *add:
		err := service.Add(*title, *desc)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Tugas ditambahkan.")
		}
	case *list:
		tasks := service.List(*search)
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tJudul\tDeskripsi\tSelesai\tTanggal Dibuat")
		for _, t := range tasks {
			fmt.Fprintf(w, "%d\t%s\t%s\t%v\t%s\n", t.ID, t.Title, t.Description, t.Completed, t.CreatedAt.Format("2006-01-02"))
		}
		w.Flush()
	case *done > 0:
		err := service.MarkDone(*done)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Tugas ditandai selesai.")
		}
	case *delete > 0:
		err := service.Delete(*delete)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Tugas dihapus.")
		}
	default:
		fmt.Println("Gunakan --help untuk melihat opsi.")
	}
}
