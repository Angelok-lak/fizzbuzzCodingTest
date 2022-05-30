# Fizzbuzz

Développer un serveur web exposant 2 routes:

- /fizzbuzz
- /stats

---

## Énoncé de l'exercice

Exercise: Write a simple fizz-buzz REST server.

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by fizz, all multiples of 5 by buzz, and all multiples of 15 by fizzbuzz. The output would look like this: 1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,....

Your goal is to implement a web server that will expose a REST API endpoint that:

Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.
Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.
The server needs to be:

Ready for production
Easy to maintain by other developers
Bonus: add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:

Accept no parameter
Return the parameters corresponding to the most used request, as well as the number of hits for this request

---

Le rendu ne contient que les packets du standard ainsi que des packets utiles à l'implémentations de tests unitaires ou encore le paquet "mux" qui est utilisé pour designé le serveur.

---

## Lancer le projet:

```
$> go run main.go
```

Afin de tester le serveur et les différentes routes exposées il suffit de faire une requête de type GET sur l'adresse du serveur avec le port associé

-> `http://localhost:8000`

Il vous suffira ensuite d'ajouter les suffixes des routes avec les paramètres associés pour pouvoir requéter
<br />
<br />

- ### `/fizzbuzz`

#### Requête:

Prend en paramètre la structure suivante dans la section body de la requête.

Tout les paramètres lister sont obligatoires dans la requête.

```
{
	"Int1"  string
	"Int2"  string
	"Str1"  string
	"Str2"  string
	"Limit" string
}
```

#### Réponse:

Renvoie le resultat du fizzbuzz avec ces paramètres ou une erreur

```
STATUS 200 OK

["1","2","3","4","bda","abd","7","8","9","bda","11","abd","13","14","bda","16"]
```

```
STATUS 400 BAD REQUEST

str2 can't be empty
```

<br />

- ### `/stats`

Ne prend aucun paramètre et renvoie la requète la plus utilisé ainsi que le nombre de fois ou celle ci a été appellé.

#### Réponse:

```
STATUS 200 OK

{
    "request": {
        "int1": "",
        "int2": "",
        "str1": "abd",
        "str2": "bda",
        "limit": ""
    },
    "hits": 3
}
```

<br />
<br />

## Tests

---

Afin de garantir l'efficacité et la qualité du code j'ai décider d'implémenter des tests unitaires sur toute la logique de gestion des entrées de la fonctions fizzbuzz afin d'en tester tout les cas, le traitement avec des paramètres valide étant toujours identique en fonction paramètres.

Les tests réalisés se trouvent dans le fichier `fizzbuzz/fizzbuzz_test.go`
<br />
<br />

### Contenu du dépôt:

---

```
fizzbuzz
│
│   README.md
│   logs.txt
│   go.mod
│   go.sum
│   main.go
│
└───fizzbuzz
│   │   error.go
│   │   fizzbuzz.go
│   │   fizzbuzz_test.go
│
└───handler
│   │   fizzbuzz.go
│   │   stats.go
│
└───logger
│   │   logger.go
│
└───stats
│   │   stats.go
```
