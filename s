[1mdiff --git a/componentes/banco_processador.go b/componentes/banco_processador.go[m
[1mindex 3f7f9c7..818e692 100644[m
[1m--- a/componentes/banco_processador.go[m
[1m+++ b/componentes/banco_processador.go[m
[36m@@ -2,7 +2,6 @@[m [mpackage componentes[m
 [m
 import ([m
 	"MESI/constantes"[m
[31m-	"fmt"[m
 )[m
 [m
 type BancoProcessadores struct {[m
[36m@@ -69,9 +68,7 @@[m [mfunc (bp *BancoProcessadores) Atualiza_Shared_Exclusive(linha int, cache_id int)[m
 [m
 		if cache.id_processador != cache_id {[m
 			linha_cache := cache.Procura_Cache(linha)[m
[31m-			fmt.Print(linha_cache)[m
 			if linha_cache != nil {[m
[31m-				fmt.Printf("NEM ENTREI AQ MANEW")[m
 				if linha_existente != nil {[m
 					println("mais de um")[m
 					return[m
[1mdiff --git a/main.go b/main.go[m
[1mindex 7e81db5..a36b75a 100644[m
[1m--- a/main.go[m
[1m+++ b/main.go[m
[36m@@ -10,11 +10,12 @@[m [mimport ([m
 [m
 func main() {[m
 	fmt.Println("Bem vindo à Biblioteca Matheus Meu Deus!")[m
[31m-	// fmt.Println("Carregando Memória Principal")[m
[32m+[m	[32mfmt.Println("Carregando Memória Principal")[m
 	mp := componentes.PreencherLivros()[m
 	componentes.Printar_MP(mp)[m
 	bp := componentes.InicializaBP(constantes.QUANTIDADE_USUARIOS)[m
 [m
[32m+[m	[32m//Loop para verificar se ele quer realizar mais uma operação ou sair do sistema.[m
 	for !should_exit() {[m
 		cache_num := escolher_cache()[m
 [m
