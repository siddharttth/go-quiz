# 🎉 Quiz Game in Go 🎉

Welcome to the Quiz Game project! This is a fun and interactive quiz game built with Go (Golang) that uses a CSV file to store questions and answers. The game features a beautiful and responsive UI with animated backgrounds and a certificate-style scorecard.

## 🚀 Features

- **Interactive Quiz**: Answer questions and see your score.
- **Beautiful UI**: Visually appealing design with animations and gradients.
- **Responsive Design**: Works well on both desktop and mobile devices.
- **Certificate**: Receive a certificate of completion with your score.

## 🛠️ Setup

Follow these steps to set up and run the project on your local machine.

### Prerequisites

- Go (Golang) installed on your machine. You can download it from [golang.org](https://golang.org/).

### Installation

1. **Clone the repository**:

    ```sh
    git clone https://github.com/yourusername/quiz-game.git
    cd quiz-game
    ```

2. **Create `questions.csv`**:

    Ensure you have a `questions.csv` file in the project root directory with the following content:

    ```csv
    question,answer
    what's my favorite food,butter chicken
    what's my favorite sport,cricket
    what's my favorite game,call of duty
    ```

3. **Run the application**:

    ```sh
    go run main.go
    ```

4. **Open your browser**:

    Navigate to `http://localhost:8080` to start the quiz game.

## 📁 File Structure
```plaintext
quiz-game/
│
├── main.go
├── questions.csv
├── static/
│   └── style.css
├── templates/
│   ├── index.html
│   └── quiz.html



## 🎨 UI and Styling

- The UI features a gradient background animation and smooth element transitions.
- The scorecard is styled like a certificate and includes the creator's name and contact information.

## 🤝 Contributing

Contributions are welcome! If you have any ideas or improvements, feel free to open an issue or submit a pull request.

## 📧 Contact

- **Creator**: Siddharth
- **Email**: siddharth.shekharr@gmail.com

Enjoy the game and have fun! 🎉

