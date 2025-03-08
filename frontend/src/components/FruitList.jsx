// src/components/FruitList.js
import React, { useEffect, useState } from 'react';
import './FruitCard.css';
import Search from './Search/Search';

const FruitList = () => {
    const [fruits, setFruits] = useState([]);
    const [error, setError] = useState(null);
    const [searchQuery, setSearchQuery] = useState('');

    const fetchAllFruits = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/fruits/all');
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            console.log('Fetched all fruits:', data);
            setFruits(data);
        } catch (error) {
            console.error('Error fetching all fruits:', error);
            setError(error.message);
        }
    };

    const fetchFruitsByName = async (name) => {
        try {
            const response = await fetch(`http://localhost:8080/api/fruits/name?name=${name}`);
            if (!response.ok) {
                const errorMessage = await response.text();
                throw new Error(errorMessage);
            }
            const data = await response.json();
            console.log('Fetched fruits by name:', data);

            // Проверяем, является ли data массивом или объектом
            let fruitsData = Array.isArray(data) ? data : [data];
            setFruits(fruitsData);
        } catch (error) {
            console.error('Error fetching fruits by name:', error);
            setError(error.message);
        }
    };

    useEffect(() => {
        fetchAllFruits();
    }, []);

    const handleAllFruitsClick = () => {
        fetchAllFruits();
    };

    const handleSearchByName = (event) => {
        event.preventDefault();
        if (searchQuery.trim() !== '') {
            fetchFruitsByName(searchQuery);
        } else {
            fetchAllFruits();
        }
    };

    return (
        <div>
            <Search
                onAllFruitsClick={handleAllFruitsClick}
                onSearchByName={handleSearchByName}
                searchQuery={searchQuery}
                setSearchQuery={setSearchQuery}
            />
            {error && <div>Error: {error}</div>}
            {fruits.length === 0 ? (
                <div>No fruits found or loading...</div>
            ) : (
                <div className="fruit-card-container">
                    {fruits.map((fruit) => (
                        <div key={fruit.name} className="fruit-card">
                            <div className="fruit-image"></div> {/* Место для изображения */}
                            <h3>{fruit.name}</h3>
                            <p>Семейство: {fruit.family}</p>
                            <p>Отряд: {fruit.order}</p>
                            <p>Род: {fruit.genus}</p>
                            <div className="nutritions-title">Витамины:</div>
                            <ul>
                                <li>Калории: {fruit.nutritions.calories}</li>
                                <li>Жиры: {fruit.nutritions.fat}g</li>
                                <li>Сахар: {fruit.nutritions.sugar}g</li>
                                <li>Углеводы: {fruit.nutritions.carbohydrates}g</li>
                                <li>Белки: {fruit.nutritions.protein}g</li>
                            </ul>
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
};

export default FruitList;