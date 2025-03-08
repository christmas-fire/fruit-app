// src/components/Search/Search.js
import React from 'react';
import './Search.css';

function Search({ onAllFruitsClick, onSearchByName, searchQuery, setSearchQuery }) {
    return (
        <div className="search">
            <div className="search-title">
                <h2>Поиск и фильтр</h2>
                <form onSubmit={onSearchByName}>
                    <input
                        type="text"
                        placeholder="Введите название"
                        value={searchQuery}
                        onChange={(e) => setSearchQuery(e.target.value)}
                    />
                    <button type="submit">Поиск по названию</button>
                </form>
                <button onClick={onAllFruitsClick}>Все фрукты</button>
            </div>
        </div>
    );
}

export default Search;