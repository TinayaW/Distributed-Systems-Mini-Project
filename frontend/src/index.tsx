import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { CssBaseline, ThemeProvider } from '@mui/material';
import { defaultTheme } from './themes/default';
import { BrowserRouter } from 'react-router-dom';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <ThemeProvider theme={defaultTheme}>
        <BrowserRouter>
          <CssBaseline />
          <App />
        </BrowserRouter>
    </ThemeProvider>
);
