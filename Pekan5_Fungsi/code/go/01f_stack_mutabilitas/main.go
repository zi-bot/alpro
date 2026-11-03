// 01f — Call Stack, Mutabilitas, dan Pure Function (versi Go).
// Jalankan: go run ./01f_stack_mutabilitas
//
// Demonstrasi konsep fungsi tingkat lanjut:
// 1. Call Stack & Stack Frames (Activation Records) melalui rekursi
// 2. Evaluasi Parameter (Pass-by-value vs Slice/Pointer Mutation):
//    - Go: Pass-by-value murni (nilai variabel disalin utuh; mutasi butuh pointer *T).
//    - Rust: Ownership & Borrow Pointer (& untuk baca, &mut untuk izin mutasi ketat).
//    - Python: Call-by-sharing (referensi objek diteruskan; mutabilitas menentukan efek).
// 3. Pure Function vs Efek Samping (Side Effects)

package main

import "fmt"

func contohStackFrame(tingkat int, kedalaman int) {
	fmt.Printf("  [Frame tingkat %d] fungsi: contohStackFrame, kedalaman stack: %d\n", tingkat, kedalaman)
	if tingkat > 1 {
		contohStackFrame(tingkat-1, kedalaman+1)
	}
	fmt.Printf("  [Pop tingkat %d] selesai dan kembali.\n", tingkat)
}

func ubahImmutable(x int) {
	// Di Go, nilai integer disalin (pass-by-value); variabel pemanggil tidak berubah
	x += 100
	fmt.Printf("    Di dalam ubahImmutable: x = %d\n", x)
}

func ubahMutable(daftar *[]int) {
	// Melewatkan pointer ke slice memungkinkan modifikasi langsung pada variabel pemanggil
	*daftar = append(*daftar, 999)
	fmt.Printf("    Di dalam ubahMutable: daftar = %v\n", *daftar)
}

func hitungDiskonMurni(subtotal int) int {
	// Pure Function: Deterministik, tanpa I/O, tanpa efek samping
	if subtotal < 100000 {
		return 0
	} else if subtotal < 250000 {
		return subtotal * 5 / 100
	} else if subtotal < 500000 {
		return subtotal * 10 / 100
	}
	return subtotal * 15 / 100
}

func main() {
	fmt.Println("=== 1. DEMONSTRASI CALL STACK & STACK FRAMES ===")
	contohStackFrame(3, 3)

	fmt.Println("\n=== 2. PASS-BY-VALUE VS POINTER/SLICE MUTATION ===")
	angka := 50
	fmt.Printf("Sebelum ubahImmutable: angka = %d\n", angka)
	ubahImmutable(angka)
	fmt.Printf("Setelah ubahImmutable : angka = %d (TIDAK BERUBAH)\n", angka)

	keranjangItem := []int{15000, 5000}
	fmt.Printf("Sebelum ubahMutable  : keranjangItem = %v\n", keranjangItem)
	ubahMutable(&keranjangItem)
	fmt.Printf("Setelah ubahMutable   : keranjangItem = %v (BERUBAH: side effect!)\n", keranjangItem)

	fmt.Println("\n=== 3. PURE FUNCTION (DETERMINISTIK & TANPA SIDE EFFECT) ===")
	subtotal := 200000
	d1 := hitungDiskonMurni(subtotal)
	d2 := hitungDiskonMurni(subtotal)
	fmt.Printf("Panggilan 1 hitungDiskonMurni(%d) = %d\n", subtotal, d1)
	fmt.Printf("Panggilan 2 hitungDiskonMurni(%d) = %d (Selalu konsisten, bebas galat)\n", subtotal, d2)
}
