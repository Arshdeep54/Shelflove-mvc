-- MySQL dump 10.13  Distrib 8.0.37, for Linux (x86_64)
--
-- Host: localhost    Database: shelflove
-- ------------------------------------------------------
-- Server version	8.0.37-0ubuntu0.22.04.3

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `TEST`
--

DROP TABLE IF EXISTS `TEST`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `TEST` (
  `ISIT` tinyint(1) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TEST`
--

LOCK TABLES `TEST` WRITE;
/*!40000 ALTER TABLE `TEST` DISABLE KEYS */;
INSERT INTO `TEST` VALUES (0);
/*!40000 ALTER TABLE `TEST` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `book`
--

DROP TABLE IF EXISTS `book`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `book` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `publication_date` date DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  `genre` varchar(255) DEFAULT NULL,
  `description` longtext,
  `rating` float DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=278 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `book`
--

LOCK TABLES `book` WRITE;
/*!40000 ALTER TABLE `book` DISABLE KEYS */;
INSERT INTO `book` VALUES (233,'The Lord of the Rings','J. R. R. Tolkien','1954-12-14',2,'Fantasy','The Lord of the Rings is an epic 1  high fantasy novel a  by the English author and scholar J. R. R. Tolkien. Set in Middle-earth, the story began as a sequel to Tolkien   s 1937 children   s book The Hobbit, but eventually developed into a much larger work. Written in stages between 1937 and 1949, The Lord of the Rings is one of the best-selling books ever written, with over 150 million copies sold. 2 ',9.6,'floor 1, rack 20, shelf 3'),(234,'Pride and Prejudice','Jane Austen','1813-01-28',8,'Romance','Pride and Prejudice is the second novel by English author Jane Austen, published in 1813. A novel of manners, it follows the character development of Elizabeth Bennet, the protagonist of the book, who learns about the repercussions of hasty judgments and comes to appreciate the difference between superficial goodness and actual goodness.',4.3,'floor 3, rack 10, shelf 2'),(235,'To Kill a Mockingbird','Harper Lee','1960-07-11',3,'Drama','To Kill a Mockingbird is a novel by the American author Harper Lee. It was published in June 1960 and became instantly successful. In the United States, it is widely read in high schools and middle schools. To Kill a Mockingbird has become a classic of modern American literature; a year after its release, it won the Pulitzer Prize. The plot and characters are loosely based on Lees observations of her family, her neighbors and an event that occurred near her hometown of Monroeville, Alabama, in 1936, when she was ten.',4.7,'floor 1, rack 25, shelf 4'),(236,'The Great Gatsby','F. Scott Fitzgerald','1925-04-10',1,'Fiction','The Great Gatsby is a 1925 novel by American writer F. Scott Fitzgerald. Set in the Jazz Age on Long Island, near New York City, the novel depicts first-person narrator Nick Carraway   s interactions with mysterious millionaire Jay Gatsby and Gatsby   s obsession to reunite with his former lover, Daisy Buchanan.',4.1,'floor 2, rack 12, shelf 3'),(237,'The Catcher in the Rye','J. D. Salinger','1951-07-16',5,'Bildungsroman novel','The Catcher in the Rye is a novel by American author J. D. Salinger that was partially published in serial form in 1945-46 before being novelized in 1951. Originally intended for adults, it is often read by adolescents for its themes of angst and alienation, and as a critique of superficiality in society. 4  5  The novel also deals with themes of innocence, identity, belonging, loss, connection, sex, and depression. The main character, Holden Caulfield, has become an icon for teenage rebellion. 6  Caulfield, nearly of age, gives his opinion on a wide variety of topics as he narrates his recent life events.',4.2,'floor 3, rack 8, shelf 1'),(238,'The Hitchhiker   s Guide to the Galaxy','Douglas Adams','1979-10-12',7,'Science Fiction','The Hitchhikers Guide to the Galaxy a  b  is a comedy science fiction franchise created by Douglas Adams. Originally a 1978 radio comedy broadcast on BBC Radio 4, it was later adapted to other formats, including novels, stage shows, comic books, a 1981 TV series, a 1984 text adventure game, and 2005 feature film.',4.2,'floor 2, rack 15, shelf 1'),(239,'One Hundred Years of Solitude','Gabriel García Márquez','1967-05-30',4,'Magical Realism','One Hundred Years of Solitude (Spanish: Cien años de soledad, Latin American Spanish:  sjen ˈaɲos ðe soleˈðað ) is a 1967 novel by Colombian author Gabriel García Márquez that tells the multi-generational story of the Buendía family, whose patriarch, José Arcadio Buendía, founded the fictitious town of Macondo. The novel is often cited as one of the supreme achievements in world literature',4.8,'floor 4, rack 5, shelf 2'),(240,'Frankenstein','Mary Shelley','1818-01-01',2,'Gothic Fiction','Frankenstein; or, The Modern Prometheus is an 1818 novel written by English author Mary Shelley. Frankenstein tells the story of Victor Frankenstein, a young scientist who creates a sapient creature in an unorthodox scientific experiment. Shelley started writing the story when she was 18, and the first edition was published anonymously in London on 1 January 1818, when she was 20. Her name first appeared in the second edition, which was published in Paris in 1821.',4.1,'floor 1, rack 18, shelf 4'),(241,'The Adventures of Huckleberry Finn','Mark Twain','1885-12-18',6,'Adventure','Adventures of Huckleberry Finn is a novel by American author Mark Twain, which was first published in the United Kingdom in December 1884 and in the United States in February 1885. \n Commonly named among the Great American Novels, the work is among the first in major American literature to be written throughout in vernacular English, characterized by local color regionalism. It is told in the first person by Huckleberry \"Huck\" Finn, the narrator of two other Twain novels (Tom Sawyer Abroad and Tom Sawyer, Detective) and a friend of Tom Sawyer. It is a direct sequel to The Adventures of Tom Sawyer.',4.4,'floor 2, rack 10, shelf 3'),(242,'1984','George Orwell','1949-06-08',6,'Dystopian Fiction','Nineteen Eighty-Four (also published as 1984) is a dystopian novel and cautionary tale by English writer George Orwell. It was published on 8 June 1949 by Secker & Warburg as Orwell   s ninth and final book completed in his lifetime. Thematically, it centres on the consequences of totalitarianism, mass surveillance, and repressive regimentation of people and behaviours within society. Orwell, a staunch believer in democratic socialism and member of the anti-Stalinist Left, modelled the Britain under authoritarian socialism in the novel on the Soviet Union in the era of Stalinism and on the very similar practices of both censorship and propaganda in Nazi Germany. 5  More broadly, the novel examines the role of truth and facts within societies and the ways in which they can be manipulated./',4.7,'floor 3, rack 1, shelf 1'),(243,'The Handmaid   s Tale','Margaret Atwood','1985-04-26',5,'Dystopian Fiction','The Handmaid   s Tale is a futuristic dystopian novel[6] by Canadian author Margaret Atwood published in 1985.[7] It is set in a near-future New England in a patriarchal, totalitarian theonomic state known as the Republic of Gilead, which has overthrown the United States government.[8] Offred is the central character and narrator and one of the   Handmaids  : women who are forcibly assigned to produce children for the   Commanders  , who are the ruling class in Gilead.',4.5,'floor 4, rack 2, shelf 4'),(244,'The Lord of the Flies','William Golding','1954-09-17',3,'Allegorical Novel','Lord of the Flies is the 1954 debut novel of British author William Golding. The plot concerns a group of British boys who are stranded on an uninhabited island and their disastrous attempts to govern themselves. The novel   s themes include morality, leadership, and the tension between civility and chaos.',4.2,'floor 2, rack 8, shelf 2'),(245,'The Metamorphosis','Franz Kafka','1915-10-27',1,'Short Story','The Metamorphosis (German: Die Verwandlung) is a novel by Franz Kafka published in 1915. Its title has also been translated as The Transformation.[1] One of Kafka   s best-known works, The Metamorphosis tells the story of salesman Gregor Samsa, who wakes one morning to find himself inexplicably transformed into a huge insect German: ungeheueres Ungeziefer, lit.   monstrous vermin   and struggles to adjust to this condition. The novella has been widely discussed among literary critics, who have offered varied interpretations. In popular culture and adaptations of the novella, the insect is commonly depicted as a cockroach.',4.3,'floor 1, rack 19, shelf 1'),(246,'Things Fall Apart','Chinua Achebe','1958-11-17',8,'Historical Fiction','Things Fall Apart is the debut novel of Nigerian author Chinua Achebe, first published in 1958.[1] It depicts the events of pre-colonial life in Igboland, a cultural area in modern-day southeastern Nigeria, and the subsequent appearance of European missionaries and colonial forces in the late 19th century. It is seen as an archetypal modern African novel in English, and one of the first such novels to receive global critical acclaim. It is a staple book in schools throughout Africa and is widely studied in English-speaking countries around the world. The novel was first published in the United Kingdom in 1958 by William Heinemann Ltd and became the first work published in Heinemann   s African Writers Series.',4.6,'floor 4, rack 3, shelf 2'),(247,'Jane Eyre','Charlotte Brontë','1847-10-16',4,'Gothic Novel','Jane Eyre (AIR; originally published as Jane Eyre: An Autobiography) is a novel by the English writer Charlotte Bronte. It was published under her pen name   Currer Bell   on 19 October 1847 by Smith, Elder & Co. of London. The first American edition was published the following year by Harper & Brothers of New York.[2] Jane Eyre is a bildungsroman that follows the experiences of its eponymous heroine, including her growth to adulthood and her love for Mr Rochester, the brooding master of Thornfield Hall.',4.4,'floor 3, rack 7, shelf 3'),(248,'The Lord of the Rings','J. R. R. Tolkien','1954-12-14',2,'Fantasy','The Lord of the Rings is an epic 1  high fantasy novel a  by the English author and scholar J. R. R. Tolkien. Set in Middle-earth, the story began as a sequel to Tolkien   s 1937 children   s book The Hobbit, but eventually developed into a much larger work. Written in stages between 1937 and 1949, The Lord of the Rings is one of the best-selling books ever written, with over 150 million copies sold. 2 ',9.6,'floor 1, rack 20, shelf 3'),(249,'Pride and Prejudice','Jane Austen','1813-01-28',8,'Romance','Pride and Prejudice is the second novel by English author Jane Austen, published in 1813. A novel of manners, it follows the character development of Elizabeth Bennet, the protagonist of the book, who learns about the repercussions of hasty judgments and comes to appreciate the difference between superficial goodness and actual goodness.',4.3,'floor 3, rack 10, shelf 2'),(250,'To Kill a Mockingbird','Harper Lee','1960-07-11',3,'Drama','To Kill a Mockingbird is a novel by the American author Harper Lee. It was published in June 1960 and became instantly successful. In the United States, it is widely read in high schools and middle schools. To Kill a Mockingbird has become a classic of modern American literature; a year after its release, it won the Pulitzer Prize. The plot and characters are loosely based on Lees observations of her family, her neighbors and an event that occurred near her hometown of Monroeville, Alabama, in 1936, when she was ten.',4.7,'floor 1, rack 25, shelf 4'),(251,'The Great Gatsby','F. Scott Fitzgerald','1925-04-10',1,'Fiction','The Great Gatsby is a 1925 novel by American writer F. Scott Fitzgerald. Set in the Jazz Age on Long Island, near New York City, the novel depicts first-person narrator Nick Carraway   s interactions with mysterious millionaire Jay Gatsby and Gatsby   s obsession to reunite with his former lover, Daisy Buchanan.',4.1,'floor 2, rack 12, shelf 3'),(252,'The Catcher in the Rye','J. D. Salinger','1951-07-16',5,'Bildungsroman novel','The Catcher in the Rye is a novel by American author J. D. Salinger that was partially published in serial form in 1945-46 before being novelized in 1951. Originally intended for adults, it is often read by adolescents for its themes of angst and alienation, and as a critique of superficiality in society. 4  5  The novel also deals with themes of innocence, identity, belonging, loss, connection, sex, and depression. The main character, Holden Caulfield, has become an icon for teenage rebellion. 6  Caulfield, nearly of age, gives his opinion on a wide variety of topics as he narrates his recent life events.',4.2,'floor 3, rack 8, shelf 1'),(253,'The Hitchhiker   s Guide to the Galaxy','Douglas Adams','1979-10-12',7,'Science Fiction','The Hitchhikers Guide to the Galaxy a  b  is a comedy science fiction franchise created by Douglas Adams. Originally a 1978 radio comedy broadcast on BBC Radio 4, it was later adapted to other formats, including novels, stage shows, comic books, a 1981 TV series, a 1984 text adventure game, and 2005 feature film.',4.2,'floor 2, rack 15, shelf 1'),(254,'One Hundred Years of Solitude','Gabriel García Márquez','1967-05-30',4,'Magical Realism','One Hundred Years of Solitude (Spanish: Cien años de soledad, Latin American Spanish:  sjen ˈaɲos ðe soleˈðað ) is a 1967 novel by Colombian author Gabriel García Márquez that tells the multi-generational story of the Buendía family, whose patriarch, José Arcadio Buendía, founded the fictitious town of Macondo. The novel is often cited as one of the supreme achievements in world literature',4.8,'floor 4, rack 5, shelf 2'),(255,'Frankenstein','Mary Shelley','1818-01-01',2,'Gothic Fiction','Frankenstein; or, The Modern Prometheus is an 1818 novel written by English author Mary Shelley. Frankenstein tells the story of Victor Frankenstein, a young scientist who creates a sapient creature in an unorthodox scientific experiment. Shelley started writing the story when she was 18, and the first edition was published anonymously in London on 1 January 1818, when she was 20. Her name first appeared in the second edition, which was published in Paris in 1821.',4.1,'floor 1, rack 18, shelf 4'),(256,'The Adventures of Huckleberry Finn','Mark Twain','1885-12-18',6,'Adventure','Adventures of Huckleberry Finn is a novel by American author Mark Twain, which was first published in the United Kingdom in December 1884 and in the United States in February 1885. \n Commonly named among the Great American Novels, the work is among the first in major American literature to be written throughout in vernacular English, characterized by local color regionalism. It is told in the first person by Huckleberry \"Huck\" Finn, the narrator of two other Twain novels (Tom Sawyer Abroad and Tom Sawyer, Detective) and a friend of Tom Sawyer. It is a direct sequel to The Adventures of Tom Sawyer.',4.4,'floor 2, rack 10, shelf 3'),(257,'1984','George Orwell','1949-06-08',6,'Dystopian Fiction','Nineteen Eighty-Four (also published as 1984) is a dystopian novel and cautionary tale by English writer George Orwell. It was published on 8 June 1949 by Secker & Warburg as Orwell   s ninth and final book completed in his lifetime. Thematically, it centres on the consequences of totalitarianism, mass surveillance, and repressive regimentation of people and behaviours within society. Orwell, a staunch believer in democratic socialism and member of the anti-Stalinist Left, modelled the Britain under authoritarian socialism in the novel on the Soviet Union in the era of Stalinism and on the very similar practices of both censorship and propaganda in Nazi Germany. 5  More broadly, the novel examines the role of truth and facts within societies and the ways in which they can be manipulated./',4.7,'floor 3, rack 1, shelf 1'),(258,'The Handmaid   s Tale','Margaret Atwood','1985-04-26',5,'Dystopian Fiction','The Handmaid   s Tale is a futuristic dystopian novel[6] by Canadian author Margaret Atwood published in 1985.[7] It is set in a near-future New England in a patriarchal, totalitarian theonomic state known as the Republic of Gilead, which has overthrown the United States government.[8] Offred is the central character and narrator and one of the   Handmaids  : women who are forcibly assigned to produce children for the   Commanders  , who are the ruling class in Gilead.',4.5,'floor 4, rack 2, shelf 4'),(259,'The Lord of the Flies','William Golding','1954-09-17',3,'Allegorical Novel','Lord of the Flies is the 1954 debut novel of British author William Golding. The plot concerns a group of British boys who are stranded on an uninhabited island and their disastrous attempts to govern themselves. The novel   s themes include morality, leadership, and the tension between civility and chaos.',4.2,'floor 2, rack 8, shelf 2'),(260,'The Metamorphosis','Franz Kafka','1915-10-27',1,'Short Story','The Metamorphosis (German: Die Verwandlung) is a novel by Franz Kafka published in 1915. Its title has also been translated as The Transformation.[1] One of Kafka   s best-known works, The Metamorphosis tells the story of salesman Gregor Samsa, who wakes one morning to find himself inexplicably transformed into a huge insect German: ungeheueres Ungeziefer, lit.   monstrous vermin   and struggles to adjust to this condition. The novella has been widely discussed among literary critics, who have offered varied interpretations. In popular culture and adaptations of the novella, the insect is commonly depicted as a cockroach.',4.3,'floor 1, rack 19, shelf 1'),(261,'Things Fall Apart','Chinua Achebe','1958-11-17',8,'Historical Fiction','Things Fall Apart is the debut novel of Nigerian author Chinua Achebe, first published in 1958.[1] It depicts the events of pre-colonial life in Igboland, a cultural area in modern-day southeastern Nigeria, and the subsequent appearance of European missionaries and colonial forces in the late 19th century. It is seen as an archetypal modern African novel in English, and one of the first such novels to receive global critical acclaim. It is a staple book in schools throughout Africa and is widely studied in English-speaking countries around the world. The novel was first published in the United Kingdom in 1958 by William Heinemann Ltd and became the first work published in Heinemann   s African Writers Series.',4.6,'floor 4, rack 3, shelf 2'),(262,'Jane Eyre','Charlotte Brontë','1847-10-16',4,'Gothic Novel','Jane Eyre (AIR; originally published as Jane Eyre: An Autobiography) is a novel by the English writer Charlotte Bronte. It was published under her pen name   Currer Bell   on 19 October 1847 by Smith, Elder & Co. of London. The first American edition was published the following year by Harper & Brothers of New York.[2] Jane Eyre is a bildungsroman that follows the experiences of its eponymous heroine, including her growth to adulthood and her love for Mr Rochester, the brooding master of Thornfield Hall.',4.4,'floor 3, rack 7, shelf 3'),(263,'The Lord of the Rings','J. R. R. Tolkien','1954-12-14',2,'Fantasy','The Lord of the Rings is an epic 1  high fantasy novel a  by the English author and scholar J. R. R. Tolkien. Set in Middle-earth, the story began as a sequel to Tolkien   s 1937 children   s book The Hobbit, but eventually developed into a much larger work. Written in stages between 1937 and 1949, The Lord of the Rings is one of the best-selling books ever written, with over 150 million copies sold. 2 ',9.6,'floor 1, rack 20, shelf 3'),(264,'Pride and Prejudice','Jane Austen','1813-01-28',8,'Romance','Pride and Prejudice is the second novel by English author Jane Austen, published in 1813. A novel of manners, it follows the character development of Elizabeth Bennet, the protagonist of the book, who learns about the repercussions of hasty judgments and comes to appreciate the difference between superficial goodness and actual goodness.',4.3,'floor 3, rack 10, shelf 2'),(265,'To Kill a Mockingbird','Harper Lee','1960-07-11',3,'Drama','To Kill a Mockingbird is a novel by the American author Harper Lee. It was published in June 1960 and became instantly successful. In the United States, it is widely read in high schools and middle schools. To Kill a Mockingbird has become a classic of modern American literature; a year after its release, it won the Pulitzer Prize. The plot and characters are loosely based on Lees observations of her family, her neighbors and an event that occurred near her hometown of Monroeville, Alabama, in 1936, when she was ten.',4.7,'floor 1, rack 25, shelf 4'),(266,'The Great Gatsby','F. Scott Fitzgerald','1925-04-10',1,'Fiction','The Great Gatsby is a 1925 novel by American writer F. Scott Fitzgerald. Set in the Jazz Age on Long Island, near New York City, the novel depicts first-person narrator Nick Carraway   s interactions with mysterious millionaire Jay Gatsby and Gatsby   s obsession to reunite with his former lover, Daisy Buchanan.',4.1,'floor 2, rack 12, shelf 3'),(267,'The Catcher in the Rye','J. D. Salinger','1951-07-16',5,'Bildungsroman novel','The Catcher in the Rye is a novel by American author J. D. Salinger that was partially published in serial form in 1945-46 before being novelized in 1951. Originally intended for adults, it is often read by adolescents for its themes of angst and alienation, and as a critique of superficiality in society. 4  5  The novel also deals with themes of innocence, identity, belonging, loss, connection, sex, and depression. The main character, Holden Caulfield, has become an icon for teenage rebellion. 6  Caulfield, nearly of age, gives his opinion on a wide variety of topics as he narrates his recent life events.',4.2,'floor 3, rack 8, shelf 1'),(268,'The Hitchhiker   s Guide to the Galaxy','Douglas Adams','1979-10-12',7,'Science Fiction','The Hitchhikers Guide to the Galaxy a  b  is a comedy science fiction franchise created by Douglas Adams. Originally a 1978 radio comedy broadcast on BBC Radio 4, it was later adapted to other formats, including novels, stage shows, comic books, a 1981 TV series, a 1984 text adventure game, and 2005 feature film.',4.2,'floor 2, rack 15, shelf 1'),(269,'One Hundred Years of Solitude','Gabriel García Márquez','1967-05-30',4,'Magical Realism','One Hundred Years of Solitude (Spanish: Cien años de soledad, Latin American Spanish:  sjen ˈaɲos ðe soleˈðað ) is a 1967 novel by Colombian author Gabriel García Márquez that tells the multi-generational story of the Buendía family, whose patriarch, José Arcadio Buendía, founded the fictitious town of Macondo. The novel is often cited as one of the supreme achievements in world literature',4.8,'floor 4, rack 5, shelf 2'),(270,'Frankenstein','Mary Shelley','1818-01-01',2,'Gothic Fiction','Frankenstein; or, The Modern Prometheus is an 1818 novel written by English author Mary Shelley. Frankenstein tells the story of Victor Frankenstein, a young scientist who creates a sapient creature in an unorthodox scientific experiment. Shelley started writing the story when she was 18, and the first edition was published anonymously in London on 1 January 1818, when she was 20. Her name first appeared in the second edition, which was published in Paris in 1821.',4.1,'floor 1, rack 18, shelf 4'),(271,'The Adventures of Huckleberry Finn','Mark Twain','1885-12-18',6,'Adventure','Adventures of Huckleberry Finn is a novel by American author Mark Twain, which was first published in the United Kingdom in December 1884 and in the United States in February 1885. \n Commonly named among the Great American Novels, the work is among the first in major American literature to be written throughout in vernacular English, characterized by local color regionalism. It is told in the first person by Huckleberry \"Huck\" Finn, the narrator of two other Twain novels (Tom Sawyer Abroad and Tom Sawyer, Detective) and a friend of Tom Sawyer. It is a direct sequel to The Adventures of Tom Sawyer.',4.4,'floor 2, rack 10, shelf 3'),(272,'1984','George Orwell','1949-06-08',6,'Dystopian Fiction','Nineteen Eighty-Four (also published as 1984) is a dystopian novel and cautionary tale by English writer George Orwell. It was published on 8 June 1949 by Secker & Warburg as Orwell   s ninth and final book completed in his lifetime. Thematically, it centres on the consequences of totalitarianism, mass surveillance, and repressive regimentation of people and behaviours within society. Orwell, a staunch believer in democratic socialism and member of the anti-Stalinist Left, modelled the Britain under authoritarian socialism in the novel on the Soviet Union in the era of Stalinism and on the very similar practices of both censorship and propaganda in Nazi Germany. 5  More broadly, the novel examines the role of truth and facts within societies and the ways in which they can be manipulated./',4.7,'floor 3, rack 1, shelf 1'),(273,'The Handmaid   s Tale','Margaret Atwood','1985-04-26',5,'Dystopian Fiction','The Handmaid   s Tale is a futuristic dystopian novel[6] by Canadian author Margaret Atwood published in 1985.[7] It is set in a near-future New England in a patriarchal, totalitarian theonomic state known as the Republic of Gilead, which has overthrown the United States government.[8] Offred is the central character and narrator and one of the   Handmaids  : women who are forcibly assigned to produce children for the   Commanders  , who are the ruling class in Gilead.',4.5,'floor 4, rack 2, shelf 4'),(274,'The Lord of the Flies','William Golding','1954-09-17',3,'Allegorical Novel','Lord of the Flies is the 1954 debut novel of British author William Golding. The plot concerns a group of British boys who are stranded on an uninhabited island and their disastrous attempts to govern themselves. The novel   s themes include morality, leadership, and the tension between civility and chaos.',4.2,'floor 2, rack 8, shelf 2'),(275,'The Metamorphosis','Franz Kafka','1915-10-27',1,'Short Story','The Metamorphosis (German: Die Verwandlung) is a novel by Franz Kafka published in 1915. Its title has also been translated as The Transformation.[1] One of Kafka   s best-known works, The Metamorphosis tells the story of salesman Gregor Samsa, who wakes one morning to find himself inexplicably transformed into a huge insect German: ungeheueres Ungeziefer, lit.   monstrous vermin   and struggles to adjust to this condition. The novella has been widely discussed among literary critics, who have offered varied interpretations. In popular culture and adaptations of the novella, the insect is commonly depicted as a cockroach.',4.3,'floor 1, rack 19, shelf 1'),(276,'Things Fall Apart','Chinua Achebe','1958-11-17',8,'Historical Fiction','Things Fall Apart is the debut novel of Nigerian author Chinua Achebe, first published in 1958.[1] It depicts the events of pre-colonial life in Igboland, a cultural area in modern-day southeastern Nigeria, and the subsequent appearance of European missionaries and colonial forces in the late 19th century. It is seen as an archetypal modern African novel in English, and one of the first such novels to receive global critical acclaim. It is a staple book in schools throughout Africa and is widely studied in English-speaking countries around the world. The novel was first published in the United Kingdom in 1958 by William Heinemann Ltd and became the first work published in Heinemann   s African Writers Series.',4.6,'floor 4, rack 3, shelf 2'),(277,'Jane Eyre','Charlotte Brontë','1847-10-16',4,'Gothic Novel','Jane Eyre (AIR; originally published as Jane Eyre: An Autobiography) is a novel by the English writer Charlotte Bronte. It was published under her pen name   Currer Bell   on 19 October 1847 by Smith, Elder & Co. of London. The first American edition was published the following year by Harper & Brothers of New York.[2] Jane Eyre is a bildungsroman that follows the experiences of its eponymous heroine, including her growth to adulthood and her love for Mr Rochester, the brooding master of Thornfield Hall.',4.4,'floor 3, rack 7, shelf 3');
/*!40000 ALTER TABLE `book` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `issue`
--

DROP TABLE IF EXISTS `issue`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `issue` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `book_id` int NOT NULL,
  `issue_date` date DEFAULT NULL,
  `expected_return_date` date DEFAULT NULL,
  `returned_date` date DEFAULT NULL,
  `isReturned` tinyint(1) DEFAULT '0',
  `returnRequested` tinyint(1) DEFAULT '0',
  `issueRequested` tinyint(1) DEFAULT '0',
  `fine` float DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `book_id` (`book_id`),
  CONSTRAINT `issue_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `issue_ibfk_2` FOREIGN KEY (`book_id`) REFERENCES `book` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `issue`
--

LOCK TABLES `issue` WRITE;
/*!40000 ALTER TABLE `issue` DISABLE KEYS */;
/*!40000 ALTER TABLE `issue` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `isAdmin` tinyint(1) DEFAULT '0',
  `adminRequest` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-07-06 21:57:08
