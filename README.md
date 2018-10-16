### Sentence Gen
A program for generating sentences

## How it works
It takes a corpus of sentences in stdin and parses the trigrams from it. Then it use the trigrams to generate a new scentence by randomly selecting words based on what is most probable to come next in the sentence based on the input data.

## How to install

```
go get github.com/langest/sentenceGen
go install
```

## How to run

```
cat myLargeTextFile | scentenceGen
```
