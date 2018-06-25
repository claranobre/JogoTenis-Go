package main
import (
  "fmt"
  "time"
  "math/rand"
)
/*
Varíavel para caso o jogador tenha ganho o jogo anterior, no próximo ter um ponto de vantagem
*/
const vantagem int = 1
/*
Variáveis para o Score, Game, Sets e Match dos jogadores
*/
/*Constante para verificar se o jogador vencedor tem pelo menos dois pontos a mais que o adversário*/
const Pontos int = 2
/*Pontos do jogadores 1 e 2*/
var pontosj1, pontosj2 int
/*
Um Game é uma sequência de pontos marcados por um jogador, ganha quem marcar 4 pontos no total e
pelo menos dois pontos a mais que o adversário
>> O projeto terá 1 Game
*/
const Games int = 1

var jog string
var gamesj1, gamesj2 int
var setj1, setj2 int
var resultado string
/*
Match é uma sequência de sets e sendo o vencedor o jogador que ganhar 3 de 5 sets
*/
var matchj1, matchj2 int

/*
func f(s string) {
// Sleep up to half a second
delay := time.Duration(r.Int() % 500) * time.Millisecond
time.Sleep(delay)
fmt.Println(s)
}*/
/*
Um Set é um conjunto de games, vence o jogador que ganhar pelo menos 6 games e pelo menos 2 games
a mais que o oponente
>> O projeto terá 1 Set
*/
const Sets int = 1

func set() {
  pontosj1, pontosj2 = 0, 0
  gamesj1, gamesj2 = 0 , 0
}
/*
Os dois jogadores possuem dois estados:
(i) esperando para receber a bola
(ii) mandando a bola de volta para o adversário
*/
var vencedor string

/*func estadoJogador(estado chan bool){
est := <- estado
if estado == true{
fmt.Println("Jogador mandou a bola para o adversário")
}else{
fmt.Println("Jogador esperando para receber a bola")
}
}*/

func jogador1(jogador chan string){
  for{
    jogador <- "j1"
    //estadoJogador()
  }
}

func jogador2(jogador chan string){
  for{
    jogador <- "j2"
    //estadoJogador()
  }
}

func painel(){
  fmt.Println("\n\n-------------Jogo Tenis-------------")
}

func exibirPlacar(){
  fmt.Println("\n\nPLACAR\n\n")
  fmt.Println("|Jogadores | Set 1 | Jogo | Pontos")
  fmt.Println("jogador 1 ", setj1 , gamesj1, pontosj1)
  fmt.Println("jogador 2 ", setj2, gamesj2, pontosj2)
}

func jogoTenis(jogador chan string){
  for{
    jog := <- jogador
    resultado = "perdeu"
    if rand.Float32() < 0.5{
      resultado = "venceu"
      if jog == "j1"{
        vencedor = "j1"
        pontosj1++
        }else{
          pontosj2++
          vencedor = "j2"
        }
        }else{
          if jog == "j1"{
            vencedor = "j2"
            pontosj2++
            }else{
              vencedor = "j1"
              pontosj1++
            }
          }
          calcularPontos()
          exibirPlacar()
          time.Sleep(time.Second * 2)
        }
      }

      func calcularPontos(){
          if vencedor == "j1"{
            if pontosj1 >= Pontos && pontosj1 >= pontosj2 + vantagem{
              gamesj1++
              pontosj1, pontosj2 = 0, 0
            }
            if gamesj1 >= Games && gamesj1 >= gamesj2 + vantagem{
              setj1++
              set()
            }
            if setj1 >= Sets && setj1 >= setj2 + vantagem{
              fmt.Println("Jogador 1 venceu")
              return
            }
          }

          if vencedor == "j2"{
            if pontosj2 >= Pontos && pontosj2 >= pontosj1 + vantagem{
              gamesj2++
              pontosj2, pontosj1 = 0, 0
            }
            if gamesj2 >= Games && gamesj2 >= gamesj1 + vantagem{
              setj2++
              set()
            }
            if setj2 >= Sets && setj2 >= setj1 + vantagem{
              fmt.Println("Jogador 2 venceu")
              return
            }
          }
      }

      /*O Jogo será composto de apenas um Set
      func marcarSet(){
      if game1 ==  1{
      gamesj1 := 1;
      }else if game1 == 2{
      gamesj1 := 2;
    }
    if game1 == 3{
    gamesj1 := 3;
    }else if game1 == 4{
    gamesj1 := 4;
  }
  if game1 == 5{
  gamesj1 := 5;
  }else if game1 == 6{
  gamesj1 = 6;
}
if game2 == 1{
gamesj2 := 1;
}else if game2 == 2{
gamesj2 := 2;
}
if game2 == 3{
gamesj2 := 3;
}else if game2 == 4{
gamesj2 := 4;
}
if game2 == 5{
gamesj2 := 5;
}else if game2 == 6{
gamesj2 := 6;
}*/
/*
Condição de parada do jogo de cada jogador

if (gamesj1 == 6 || gamesj2 == 6){
game2 := 0;
game1 := 0;
}
if (gamesj2 == 6 || gamesj1 == 6){
game1 := 0;
game2 := 0;
}
}*/

func main(){
  painel();

  var jogador chan string = make(chan string)

  rand.Seed(time.Now().UnixNano())

  go jogador1(jogador)
  go jogador2(jogador)

  go jogoTenis(jogador)

  /* Wait for 6 more seconds to let all go routine finish
  time.Sleep(time.Duration(6) * time.Second)
  fmt.Println("--- Done.")*/
  var input string
  fmt.Scanln(&input)
}
