# cs686_BlockChain_P2_starter

## What is the application?
A Book review application. Allows users to post new reviews and ratings of previously read books.

## The Problem
Research reveals that 93 percent of shoppers say the testimonials they read affect their purchasing decisions. 

Not all review sites, nor the reviews left on them are genuine. The trust placed in these platforms means there’s a lot of manipulation going on review sites. According to the same survey, four in five of us have read a fake review in the past year, and 84 percent of us have admitted we struggle to decipher which posts are authentic.

One of the reasons why writing reviews are so easy is that review writers don’t have to prove that they have purchased a product – and, in some cases, they don’t even need to provide their name or use an account to submit their review. 

Research done by Podium. 
https://www.google.com/search?q=podium&oq=podium+&aqs=chrome..69i57j69i60j0l4.1034j1j7&sourceid=chrome&ie=UTF-8 

## Why do we need blockchain technology for this application?
Blockchain technology allows us to satisfy both authenticity and integrity of book reviews: 

- Authenticity: Reviews posted can be verified that it came from the corresponding reviewer. 
- Integrity: Reviews posted onto the site cannot be tampered with by a third party as it has to be signed with a private key. 

We can achieve these two features using public-private key cryptology on the blockchain platform, coupled with blockchain technology we can ensure immutability via our proof of work mechanism.

## Workflow 
#### Lifecycle of a new book review 
1. Client registers with the platform by obtaining a new public and private key - GetPrivateKey();
2. Clients contructs a new ReviewObject, signs it with the private key - NewReview(review string);
3. The signed ReviewObject is broadcast to the peerlist of nearby peer nodes - GetPeerList();
4. Miner nodes have a pool of ReviewObjects which they store all incoming ReviewObjects before they are processed. 
5. Miner nodes add ReviewObjects to their block up to the size of the BlockSizeLimit;
6. Miners attempt to solve the ProofOfWork required to earn the BlockReward and transaction fees associated with the ReviewObjects. 
7. They broadcast the solved block to the peerlist of peer nodes.

#### Lifecycle of obtaining all book reviews 
1. Client submits an operation to view all book reviews with a transaction fee associated with the operation. 
2. The miner who services this operation is rewarded the transaction fee. 
3. The client is able to view all the reviews.

## What are the list of functionalities?

#### Allows a client to retrieve all book reviews that exist on the Book Review Application
```
GET /allBookReviews
```

```
GET /private-public key pair
{
"public": string,
"private": string,
}
```

#### Takes the client to the page where a new book review can be created 
```
GET /createBookReview
```

#### Allows a client to post a new book review onto the Book Review Application
##### Implementation details
```
POST /newBookReview
Content-type: application/json

{"title": string,
"reviewText": string,
"reviewRating": int,
"reviewerName": string,
"transactionFee": float,
"senderAddress": string,
"signature": string,
}
```

- MPT: 
key = timestamp
value = bookReview JSON element 
 
- title: User generated content
- reviewText: User generated content
- reviewRating: User generated content <Value from 0-5>
- transactionFee: User generated content
- signature: Generate new signature key for client (endpoint)
- senderAddress: Sender public key 
- signature: privateKey(BookReview);

#### Allows a client to generate a new book Id
```
GET /newBookId

```
#### Allows a client to generate a new key signature
```
GET /newBookId

```


## Define the success of this product?
 - [ ] Users will be able to post new reviews to the application 
 - [ ] Users will be able to view all reviewed books on the platform

## Midpoint Task List
 - [x] Set up new data structures 
 - [x] POST newBookReview endpoint
 - [x] GET allBookReviews endpoint
 - [x] GET createBookReview endpoint

## Final Deadline Task List
 - [ ] NewBookReview FrontEnd 
 - [ ] AllBookReviews FrontEnd
 - [ ] Public Private Key generation for clients
 - [ ] Workflow logic for POST /newBookReview
 - [ ] Workflow logic for GET /allBookReviews
 - [ ] Workflow logic for GET /createBookReview
 - [ ] Miner wallet to store transaction fees 

## Disclaimer 
This project is an academic level abstract of an actual decentralized application on the public blockchain. 
The implementation of the underlying blockchain is simplified to fit the nature of this project, and several assumptions have been made about transactions to allow the project to work seamlessly. 

1. Client that posts a new book review owns sufficient transaction fees 
2. New book review is an actual book
