# CoralBlockchain_Assignment
OS Used: Ubuntu 16.04,
Backend Language: Golang,
Database: MySQL,
FrontEnd: HTML, Ajax,
WebService: REST APIs.

Functionalities Implemented:
1) Insert API: Whenever this POST API is hit, the database is searched for the given email-id, since email-id is the primary key. If such a record exists, it is updated, but if not, a new record is created. The current date-time info is added in the date time field of the record.
2) Search API: Search is done by this GET API. It takes the email id as the path variable. This variable is used to search the MYSQL database for any record with such an email-id. If record exists, it is returned in the form of JSON data. Else, success as false is returned, with the message that record doesn't exists.
3) Delete API: Search is done by this DELETE API. It takes the email id as the path variable. This variable is used to search the MYSQL database for any record with such an email-id. If record exists, it is deleted. Else, success as false is returned, with the message that record doesn't exists.

Overview:
The Backend is written in Golang, and not NodeJS because I am not very familiar with NodeJS(and good at Golang). But, if a task demands NodeJS or any other language, I am always ready to learn and work on it. For instance, when I joined Cubereum(the place where I am working right now), I didn't know much about Golang. But since Golang was important to Hyperledger Fabric, and the company wished to have its backend in Golang, I learnt it in no time.
Front-End in HTML isn't something I have worked on(I have made android app frontends in XML). So, this frontend isn't very attractive or appeasing. The search and delete functionalities work, but do not display the response due to some problem.
The REST APIs are running on localhost, port 8000, working perfectly.
The MySQL database is being used to store data, and is synced with golang. Queries were written using the MySQL monitor, to create databases, tables, users and test other required queries.
In all, it was a fun assignment, and I learnt a lot of new things.
