package main
import "fmt"

func main() {
    const NMAX int = 100
    var A [NMAX]int
    var B [NMAX]int
    var i, n, target, j, count int

    fmt.Scan(&n, &target)
    for i = 0; i < n; i++ {
        fmt.Scan(&A[i])
    }

    // Cari pasangan yang jumlahnya = target
    for i = 0; i < n; i++ {
        for j = i + 1; j < n; j++ {
            if A[i]+A[j] == target {
                B[count] = A[i]
                count++
                B[count] = A[j]
                count++
                break
            }
        }
    }

    // Cetak hanya elemen yang terisi
    fmt.Print("[")
    for i = 0; i < count; i++ {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(B[i])
    }
    fmt.Println("]")
}