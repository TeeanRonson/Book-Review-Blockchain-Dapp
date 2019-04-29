# cs686_BlockChain_P2_starter

## What is the application?
An application that allows users to post new reviews and ratings of books. 

## The Problem
Research reveals that 93 percent of shoppers say the testimonials they read affect their purchasing decisions. 

Not all review sites, nor the reviews left on them are genuine. The trust placed in these platforms means there’s a lot of manipulation going out on there. According to the same survey, four in five of us have read a fake review in the past year, and 84 percent of us have admitted we struggle to decipher which posts are authentic.

One of the reasons why writing reviews are so easy is that review writers don’t have to prove that they have purchased a product – and, in some cases, they don’t even need to provide their name or use an account to submit their review. 

Research done by Podium. 
https://www.google.com/search?q=podium&oq=podium+&aqs=chrome..69i57j69i60j0l4.1034j1j7&sourceid=chrome&ie=UTF-8 

## Why do we need blockchain technology for this application?
Blockchain technology could confidentially keep track of a consumer’s purchases and only enable them to leave reviews for the products they have bought – plus prevent people from writing multiple posts. 

For example, if their shopping history is locked into the blockchain, the purchase could be checked to see whether the item someone is trying to review is something they have bought in the past.

The main advantage blockchain technology brings to this application is immutability and trust. A user that posts a review and rating to the application will have to sign it with his/her private key signature and this is locked into the blockchain. 

## What are the list of functionalities?

#### Allows a client to retrieve all book reviews that exist on the Book Review Application
Implementation details

```
GET /allBookReviews
```

#### Takes the client to the page where a new book review can be created 
Implementation details

```
GET /createBookReview
```

#### Allows a client to post a new book review onto the Book Review Application
Implementation details


```
POST /newBookReview
Content-type: application/json

{"title": string,
"bookId": int,
"description”": string,
"reviewText": string,
"reviewRating": int,
"reviewerName": string,
"transactionFee": float,
"signature": string,
}
```

## Define the success of this product?
[ ] Users will be able to post new reviews to the application 
[ ] Users will be able to view all reviewed books on the platform

## Midpoint Task List
[ ] Set up new data structures 
[ ] POST newBookReview
[ ] GET allBookReviews

## Final Deadline Task List
[ ] NewBookReview FrontEnd 
[ ] AllBookReviews FrontEnd
[ ] Public Private Key generation for clients
[ ] Workflow logic for POST /newBookReview
[ ] Workflow logic for GET /alllBooksReviews
[ ] Miner wallet to store transaction fees 

## Disclaimer 
This project is an academic level abstract of an actual decentralized application on the public blockchain. 
The implementation of the underlying blockchain is simplified to fit the nature of this project, and several assumptions have been made about transactions to allow the project to work seamlessly. 





