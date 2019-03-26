## Needed /Tech stacks
    + go
    + mysql
    
## To get started follow this checklist:
    + create schema golang
    + go get -u github.com/go-sql-driver/mysql
    + run sql
      DROP TABLE IF EXISTS `employee`;
      CREATE TABLE `employee` (
        `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
        `name` varchar(30) NOT NULL,
        `city` varchar(30) NOT NULL,
        PRIMARY KEY (`id`)
      ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
    + go run main.go
   
