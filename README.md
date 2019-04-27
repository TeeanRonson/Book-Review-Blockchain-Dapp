# cs686_BlockChain_P2_starter

## What is the application?
An application that allows users to review and rate books. 

## The problem
Research reveals that 93 percent of shoppers say the testimonials they read affect their purchasing decisions.

## Why do we need blockchain for this application?
The main advantage blockchain technology brings to this application is immutability. A user that posts a review and rating to the application will have to sign it with his/her private key signature and this is locked into the blockchain. 

It is immutable so that no actors can change their previously proposed book review and rating.

What are the list of functionalities?
Propose holiday  { “title”: string, “description”: string, “review”: string, “reviewerName”: string, “proposerId”: int, “rating”: int, “bookImage”: string }
Get all books
Top rated books
Books rated by memberName
Get all book members

Define the success of this product?
Users will be able to view the top voted books recommended on the platform 

Midpoint 
Set up new data structures 
Proposal of a new book should be hashed into a string before adding into MPT 
Set up API calls 
Set up client book review proposal 
