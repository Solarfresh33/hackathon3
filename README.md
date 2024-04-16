# <div align="center">TRACKY

## SOMMAIRE

- [I. Comment installer Tracky](#i-comment-installer-tracky)
- [II. Hebergement de Tracky](#ii-hebergement-de-tracky)
- [III. Fonctionnement de Tracky](#iii-fonctionnement-de-tracky)
- [IV. Page Administrateur](#iv-page-administrateur)
- [V. Les fonctionnalitées de Tracky](#v-les-fonctionnalitées-de-tracky)
- [VI. Si Tracky ne fonctionne pas](#vi-si-tracky-ne-fonctionne-pas)


## I. Comment installer Tracky

Pour installer Tracky, commencez par cloner le repository sur votre ordinateur en utilisant le terminal et la commande suivante :

```bash
git clone https://ytrack.learn.ynov.com/git/cmaxime/hackaton
```
Ensuite, lancez le site via les commandes suivantes:

```go
cd ./hackaton
go run ./server/main.go
```
Pour accéder au site, ouvrez un navigateur et entrez l'adresse suivante : http://localhost:8080/

## II. Hebergement de Tracky

Nous avons choisi d'héberger Tracky. Pour accéder à Tracky, veuillez vous connectez à l'adresse suivante : https://groupe5.etudiants.ynov-bordeaux.com/

## III. Fonctionnement de Tracky

Tracky est une application web qui sert a suivre vos colis partout dans la France. Si une commande a a été effectuée, un Email vous sera envoyé avec votre numéro de colis dedans. Vous pouvez donc cliquer sur le bouton "Suivre mon colis" dans l'email ou copier-coller le numéro de votre colis à ces adresses ci dessous : http://localhost:8080/ ou https://groupe5.etudiants.ynov-bordeaux.com/ 

## IV. Page Administrateur

Vous pourrez accéder a la page administrateur à cette adresse : https://groupe5.etudiants.ynov-bordeaux.com/login, contactez un des développeurs du site afin d'avoir des accès administrateur, sinon en localhost lors du premier lancer, allez a cette page : http://localhost:8080/login, tapez admin en pseudo et admin en mot de passe. Alors vous serrez connecté en tant qu'administrateur sur Tracky pour gérer les différents colis du client. 

La page scan permet de scanner un QR Code afin de modifier l'avancement du colis, la page créer permet de créer un colis lorsqu'une commande à été effectuée, la page utilisateur permet d'ajouter un administrateur au site, la page livraison permet de voir tous les colis qui ont été commandé, la page update permet de mettre à jour l'état d'avancement du colis grâce à son numéro de colis, et enfin, la page application permet de télécharger notre application.

## V. Les fonctionnalitées de Tracky
<<<<<<< HEAD
Tracky a mit en place un bot qui peut vous guidez sur le site. Il est également possible pour les administrateurs de pouvoir scanner un QR Code pour pouvoir modifier l'avancement du colis avec leur téléphone par exemple. Et enfin, il est possible de devenir affilié chez Tracky afin de se faire livrer un colis **(Julien add une explication de ta feature pour devenir livreur et tout, j'sais pas comment l'expliquer mdr)** . 
=======
Tracky a mit en place un bot qui peut vous guidez sur le site. Il est également possible pour les administrateurs de pouvoir scanner un QR Code pour pouvoir modifier l'avancement du colis avec leur téléphone par exemple. Et enfin, il y a un système de livraison premium qui se réfère a des chauffeurs indépendants a la manière de Uber. 
>>>>>>> 68824f80be145e45089d5e51c094387f0b26885d

## VI. Si Tracky ne fonctionne pas

Si Tracky ne fonctionne pas en local, suivez ces étapes :

1. Vérifiez l'installation de Go sur votre ordinateur.

2. Assurez-vous que le repository du hackaton est correctement installé.

3. Si le problème persiste, supprimez le dossier hackaton et recommencez l'installation.

4. Si aucune de ces solutions ne fonctionne, contactez les auteurs de Tracky.

## <div align="right">Les auteurs de Tracky
<div align="right">CANO Kévin
<div align="right">DUBLANC Noémie
<div align="right">PAYET Lionel
<div align="right">CHAMOULEAU Julien
<div align="right">CHORT Maxime
<div align="right">ISIDORE Maxime

