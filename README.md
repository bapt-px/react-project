# react-project

1 - installer les dependances: 
  à la racine du projet lancer "npm install"
  
2 - Lancer le serveur: 
  Ouvrir dans une console le dossier server puis lancer la commande "go run main.go"
  Celui-ci démarre sur le port 8083

3 - Démarrer l'application web: 
  Dans une autre console à la racine du projet, lancer la commande "npm run dev"
  L'application devrait s'ouvrir automatiquement, sinon, aller sur l'url http://localhost:3000
  
L'api prend en paramètre la donnée "url" puis retourne la structure de donnée suivante: { average: float64, nbComment: int }  
