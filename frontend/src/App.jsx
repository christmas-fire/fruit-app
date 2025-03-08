// src/App.js
import React from 'react';
import './App.css';
import FruitList from './components/FruitList';
import Header from './components/Header/Header';
import Footer from './components/Footer/Footer';
import Search from './components/Search/Search';

function App() {
  return (
    <div className="App">
        <Header />
        <FruitList />
        <Footer />
    </div>
  );
}

export default App;