import React from 'react';
import DefaultRouter from './routers/DefaultRouter';
import HomePage from './pages/HomePage';
import AboutPage from './pages/AboutPage';
import NotFoundPage from './pages/NotFoundPage';

const routes = [
  { path: '/', component: HomePage },
  { path: '/about', component: AboutPage },
  { path: '*', component: NotFoundPage }, 
];

function App() {
  return (
    <DefaultRouter routes={routes} />
  );
}

export default App;
