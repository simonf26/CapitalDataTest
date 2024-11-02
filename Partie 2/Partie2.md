# Partie 2

-Quels sont les noms, prénoms et adresse mails des clients :
    
    - ayant achetés au moins un article au cours de l’année 2019 ?
```sql
SELECT c.firstname, c.lastname, c.email FROM Client c 
JOIN sale s ON c.card_number = s.card_number 
WHERE YEAR(s.date)=2019;
```

    - ayant achetés au moins une bouteille de Coca au mois de Janvier 2019 ?
```sql
SELECT c.firstname, c.lastname, c.email FROM Client c 
JOIN sale s ON c.card_number = s.card_number 
JOIN product p on s.product_id = p.id
WHERE p.name LIKE "Coca";
```
-Combien de clients différents ont acheté :
    - au moins un produit de la catégorie «Chocolat» ?
```sql
SELECT COUNT (DISTINCT c.card_number) FROM Client c
JOIN sale s ON c.card_number = s.card_number 
JOIN product p ON s.product_id = p.id
JOIN category cat ON cat.category_id = p.category_id
WHERE cat.name = "Chocolat";
```
    - au moins 3 produits de la catégorie «Chocolat» ?
```sql
SELECT COUNT(DISTINCT c.card_number) FROM Client c
JOIN sale s ON c.card_number = s.card_number
JOIN product p ON s.product_id = p.id
JOIN category cat ON cat.category_id = p.category_id
WHERE cat.name = "Chocolat"
HAVING SUM(s.quantity) >=3;
```
-Quel est le chiffre d’affaire du magasin ?

    - le chiffre d’affaire sur la catégorie de produit «Produits nettoyants» ?
```sql
SELECT SUM(p.price * v.quantity) FROM sale s
JOIN product p ON s.product_id = p.id
JOIN category cat ON p.category_id = cat.id
WHERE cat.name = "Produits nettoyants";
```
-Quels sont les trois catégories de produit les plus vendus ?
    
    - les plus vendus par sexe ?    
```sql
SELECT c.gender, cat.name FROM category cat
JOIN product p ON cat.id = p.category_id
JOIN sale s ON p.id = s.product_id
JOIN client c ON s.card_number = c.card_number
GROUP BY c.gender
ORDER BY SUM(v.quantity) DESC
LIMIT 3;
```