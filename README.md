# Stairway to go

## Installare GO
Per gestire più versioni di GO sulla stessa macchina è utile utilizzare [gvm](https://github.com/moovweb/gvm).
```sh
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

Assicurati di inserire in .bashrc:
```
source ~/.gvm/scripts/gvm
```

Per ottenere la lista delle distribuzioni di go disponibili:
```
gvm listall
```

Per installarne una (l'opzione -B serve a installare go utilizzando binari precompilati):
```
gvm install go1.14.3 -B
```

Per usare una versione di go specifica:
```
gvm use go1.14.3
```

Questo comando:
- mette a disposizione il comand `go` (`which go`)
- setta le variabili d'ambiente $GOROOT e $GOPATH

Per avere una lista delle versioni disponibili nella macchina corrente:
```
gvm list
```

Per avere una versione di go di default sempre pronta nel sistema, aggiungere nel .bashrc:
```
echo "gvm use go1.14.3" >> ~/.zshrc
```

## Differenza tra $GOROOT e $GOPATH
- $GOROOT: directory dove vengono cercati binari di esecuzione
- $GOPATH: directory che contiene dipendenze esterne

## GOPATH vs go mod
Con GOPATH si gestiscono le dipendenze globali. In alternativa, per gestire dipendenze di un progetto alla `npm` è disponibile `go mod`.

**Inizializzare un progetto**
`go mod init github.com/menxit/stairway-to-go`

Questo comando creerà un file `go.mod` che specificherà nome modulo e versione go utilizzata.
```
module github.com/menxit/stairway-to-go

go 1.14
```

**Installare una dipendenza**
```
go get github.com/stretchr/testify
```

Leggiamo come è stato aggiornato go.mod:
```
module github.com/menxit/stairway-to-go

go 1.14

require github.com/stretchr/testify v1.5.1 // indirect
```

Il commento `// indirect` è stato aggiunto per indicare che la dipendenza non è ancora mai stata utilizzata nel progetto. È possibile utilizzare il comando `go mod tidy` per rimuovere tutte le dipendenze importate e non utilizzate nel progetto. Ha senso essere certi di avviare questo comando ogni volta che si rilascia una release.

Oltre al file `go.mod` viene creato un file `go.sum`. Il file `go.mod` già funge da lock, il `go.su` funge da ulteriore modifica (check di integrità delle dipendenze).

**Dove stanno le dipendenze?**
Non come in npm, niente cartella `node_modules`. Tutte le dipendenze scaricate si trovano in $GOPATH e vengono condivise tra tutti i moduli presenti sulla macchina corrente.

## Tests
I file di test vanno salvati nella forma:
```
qualchecosa_test.go
```

E vanno lanciati con il comando:
```
go test ./path/to/tests
```

Se si vuole un output verboso:
```
go test -v ./path/to/tests
```

## Tool utili
- [json to go](https://mholt.github.io/json-to-go/)
