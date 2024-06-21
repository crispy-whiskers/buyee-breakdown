# buyee-breakdown
## Buyee
So, you're using the popular international proxy for shopping on Japanese secondhand sites (Yahoo Auctions, mercari, Rakuten, etc), Buyee. The way Buyee works is you buy something off one of those sites via Buyee, and they will put their own Japanese warehouse as the recipient of the goods. As you know, international shipping is not cheap. A lot of other proxy services like Buyee will charge you that enormous shipping fee to you. However, Buyee was unique for being the first one to offer batch-ordering, a service where instead of paying shipping for every individual item you buy, you instead buy a bunch of things all at once and tell Buyee to smack it all in one box. You then pay shipping for just that box, and save loads of money.

## The Problem
Often times, you don't want enough things to justify a batch order. Maybe just one plushie and a figurine. You're not saving much money if you ship them together. You'd probably be losing money, including all the fees Buyee imposes. The solution is to get a bunch of friends together who all want to buy a couple things. This way, you have enough items to have a margin of saving via batch order. But, how do you split the final shipping cost?


## This solution
I used to do it manually on a Google Spreadsheet, with a ton of formulas to keep track of all the costs. Well, its been a few years since I started running these with my friends, and I am sick and tired of redoing the same formulas but typing in slightly adjacent cells. This project does all of this, and will spit the data back out into a spreadsheet friendly value.

### Why Golang?
Because I needed to learn it for work.

# Technologies
- Go version 1.22.4
- fyne

## Features
- Saving and reading from a JSON
- clean UI
- very, very messy code because I did this in a span of 2 days

## Installation
- Clone this repo.
- Run `go mod tidy`
- Run `go run ./src/generator-main.go`

idk how to build a target for golang yet
