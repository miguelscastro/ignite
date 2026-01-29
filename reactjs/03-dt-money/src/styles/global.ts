import { createGlobalStyle } from 'styled-components'

export const GlobalStyle = createGlobalStyle`
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;

    
  }

  :focus-visible {
      outline: 0;
      box-shadow: 0 0 0 2px ${(props) => props.theme['green-500']};
    }

    body {
      background-color: ${(props) => props.theme['gray-800']};
      color: ${(props) => props.theme['gray-100']};
      --webkit-font-smoothin: antialiased;
    }

    body, input, textarea, button {
      font: 400 1rem Roboto, sans-serif;
    }

    .screen_reader-only {
      position: absolute;
      width: 1px;
      height: 1px;
      padding: 0;
      margin: -1px;
      overflow: hidden;
      clip: rect(0, 0, 0, 0);
      white-space: nowrap;
      border: 0;
    }
`
