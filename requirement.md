#Game Application

# use case

## User use cases

### Register
user can register to application by phone number

### Login
user can log in to application by phone number and password


## Game use cases

### Each game have a given number of questions
### The difficulty level of questions are "easy, medium, hard"
### The difficulty level of questions are "easy, medium, hard"
### Game winner is determined by number of correct answers that each user answered


# Entity

## User
- ID
- Phone number
- Avatar

## Game
- ID
- Category
- Question List
- Players
- Winner

## Questions
- ID
- Question
- Answer List
- Correct Answers
- Category
- Difficulty

## Category
- ID
- Name
- Description