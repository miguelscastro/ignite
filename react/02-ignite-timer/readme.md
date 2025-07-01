# Timer

Project developed in the "Creating SPAs with ReactJS" module from the Ignite Bootcamp by Rocketseat.

**Live demo:** *https://timer-miguelscastro.vercel.app/*

## 🧠 About the project

A simple Pomodoro Timer that allows users to:

- Define the task name (cycle/workflow)
- Choose the cycle duration (in minutes)
- Start, pause, and interrupt the timer
- Log the history of completed or interrupted cycles
- View the status of each cycle (finished / interrupted)

## 🚀 Technologies

Built with:

- React (SPA)
- TypeScript
- Vite — fast bundler
- Context API + Hooks — state management
- Zod — schema validation
- date‑fns — date and time manipulation
- Immer — immutable state updates

## ⚙️ How to run

Clone the repository: `git clone https://github.com/miguelscastro/ignite.git`

enter it's folder: `cd ignite/react/02-ignite-timer`

Install dependencies: `npm install`

Run the application: `npm run dev`

Then open http://localhost:5173 in your browser.

## ✅ Features

- Create and start a cycle with a custom duration
- Pause and interrupt active cycles
- View history with real-time status (finished or interrupted)
- Persist history using localStorage
