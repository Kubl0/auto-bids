import React, { useState } from "react";
import { Link } from "react-router-dom";
import LocationInputSearch from "./LocationInputSearch";

export default function ParametersInputMain({showAllFields}: {showAllFields: boolean}) {
  const [selectedMake, setSelectedMake] = useState("");
  const [selectedModel, setSelectedModel] = useState("");
  const [locationParams, setLocationParams] = useState<{ position: [number, number] | null; radius: number }>({ position: null, radius: 100000 });

  const handleMakeChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedMake(event.target.value);
    setSelectedModel("");
  };

  const handleLocationChange = (params: { position: [number, number] | null; radius: number }) => {
    setLocationParams(params);
  };

  return (
    <div className="parameters-input-main">
      <label>Make:</label>
      <select value={selectedMake} onChange={handleMakeChange}>
        <option value="">Make</option>
        <option value="FSO">FSO</option>
        <option value="FSR">FSR</option>
        <option value="FSM">FSM</option>
      </select>

      <label>Model:</label>
      <select
        value={selectedModel}
        onChange={(event) => setSelectedModel(event.target.value)}
      >
        <option value="">Model</option>
        {selectedMake === "FSO" && (
          <>
            <option value="Polonez">Polonez</option>
            <option value="Warszawa">Warszawa</option>
          </>
        )}
        {selectedMake === "FSR" && (
          <>
            <option value="Tarpan">Tarpan</option>
          </>
        )}
        {selectedMake === "FSM" && (
          <>
            <option value="126p">126p</option>
            <option value="Syrena">Syrena</option>
          </>
        )}
      </select>

      <label>Type:</label>
      <select>
        <option value="">Type</option>
        <option value="Sedan">Sedan</option>
        <option value="Station Wagon">Station Wagon</option>
        <option value="Hatchback">Hatchback</option>
        <option value="SUV">SUV</option>
        <option value="Van">Van</option>
        <option value="Cabriolet">Cabriolet</option>
        <option value="Coupe">Coupe</option>
        <option value="Other">Other</option>
      </select>
      <label>Year:</label>
      <input type="text" placeholder="Year from" />
      <input type="text" placeholder="Year to" />
      {showAllFields &&(
      <><label>Mileage:</label>
      <input type="text" placeholder="Mileage from" />
      <input type="text" placeholder="Mileage to" />
      <label>Price:</label>
      <input type="text" placeholder="Price from" />
      <input type="text" placeholder="Price to" />
      <label>Fuel Type:</label>
      <select>
        <option value="">Fuel Type</option>
        <option value="Gasoline">Gasoline</option>
        <option value="Diesel">Diesel</option>
        <option value="LPG">LPG</option>
        <option value="Electric">Electric</option>
        <option value="Hybrid">Hybrid</option>
        <option value="Other">Other</option>
      </select>

      <label>Condition:</label>
      <select>
        <option value="">Condition</option>
        <option value="New">New</option>
        <option value="Used">Used</option>
        <option value="Damaged">Damaged</option>
      </select>
      <label>Doors:</label>
      <select>
        <option value="">Doors</option>
        <option value="2">2</option>
        <option value="3">3</option>
        <option value="4">4</option>
        <option value="5">5</option>
        <option value="6">6</option>
        <option value='Other'>Other</option>
      </select>

      <label>Engine Capacity:</label>
      <input type="text" placeholder="Engine Capacity from" />
      <input type="text" placeholder="Engine Capacity to" />

      <label>Power:</label>
      <input type="text" placeholder="Power from" />
      <input type="text" placeholder="Power to" />

      <label>Gearbox:</label>
      <select>
        <option value="">Gearbox</option>
        <option value="Manual">Manual</option>
        <option value="Automatic">Automatic</option>
        <option value="Semi-automatic">Semi-automatic</option>
        <option value="CVT">CVT</option>
        <option value="Other">Other</option>
      </select>

      <label>Drive:</label>
      <select>
        <option value="">Drive</option>
        <option value="Front">Front</option>
        <option value="Rear">Rear</option>
        <option value="All">All</option>
        <option value="Other">Other</option>
      </select>

      <label>Steering:</label>
      <select>
        <option value="">Steering</option>
        <option value="Left">Left</option>
        <option value="Right">Right</option>
      </select>

      <label>Location:</label>
      <LocationInputSearch onLocationChange={handleLocationChange} />

      <label>Status:</label>
      <select>
        <option value="">Status</option>
        <option value="New">New</option>
        <option value="Used">Used</option>
        <option value="Damaged">Damaged</option>
      </select>
      </>)
      }

      <Link to="/search">
        <button>Search</button>
      </Link>
    </div>
  );
}
