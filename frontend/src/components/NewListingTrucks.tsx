import React from "react";
import { useState } from "react";

export default function NewListingTrucks() {
    const [selectedMake, setSelectedMake] = useState("");
    const [selectedModel, setSelectedModel] = useState("");

    const handleMakeChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedMake(event.target.value);
        setSelectedModel("");
    };

    return(
        <div className="new-listing-page">
            <label>Photos:</label>
            <input type='file' multiple accept="image/*" />
            <label>Title:</label>
            <input type='text' placeholder='Title' />
            <label>Make:</label>
            <select value={selectedMake} onChange={handleMakeChange}>
                <option value="">Make</option>
                <option value="Mercedes-Benz">Mercedes-Benz</option>
                <option value="Scania">Scania</option>
                <option value="Volvo">Volvo</option>
            </select>
            <label>Model:</label>
            <select value={selectedModel} onChange={(event) => setSelectedModel(event.target.value)}>
                <option value="">Model</option>
                {selectedMake === "Mercedes-Benz" && (
                    <>
                        <option value="Actros">Actros</option>
                        <option value="Atego">Atego</option>
                        <option value="Axor">Axor</option>
                        <option value="Econic">Econic</option>
                        <option value="Unimog">Unimog</option>
                        <option value="Vario">Vario</option>
                    </>
                )}
                {selectedMake === "Scania" && (
                    <>
                        <option value="G-series">G-series</option>
                        <option value="P-series">P-series</option>
                        <option value="R-series">R-series</option>
                        <option value="S-series">S-series</option>
                    </>
                )}
                {selectedMake === "Volvo" && (
                    <>
                        <option value="FH">FH</option>
                        <option value="FM">FM</option>
                        <option value="FL">FL</option>
                        <option value="FE">FE</option>
                        <option value="FH16">FH16</option>
                    </>
                )}
            </select>
            <label>Application:</label>
            <select>
                <option value="">Application</option>
                <option value="Box">Box</option>
                <option value="Curtain Side">Curtain Side</option>
                <option value="Flatbed">Flatbed</option>
                <option value="Refrigerated">Refrigerated</option>
                <option value="Tanker">Tanker</option>
                <option value="Tipper">Tipper</option>
                <option value="Bus">Bus</option>
                <option value="Other">Other</option>
            </select>
            <label>Year:</label>
            <input type="text" placeholder="Year"/>
            <label>Mileage:</label>
            <input type="text" placeholder="Mileage"/>
            <label>Price:</label>
            <input type="text" placeholder="Price"/>
            <label>Description:</label>
            <textarea placeholder="Description" />
            <label>Condition:</label>
            <select>
                <option value="">Condition</option>
                <option value="New">New</option>
                <option value="Used">Used</option>
                <option value="Damaged">Damaged</option>
            </select>
            <label>Fuel:</label>
            <select>
                <option value="">Fuel</option>
                <option value="Diesel">Diesel</option>
                <option value="Electric">Electric</option>
                <option value="LPG">LPG</option>
                <option value="CNG">CNG</option>
                <option value="LNG">LNG</option>
                <option value="Gasoline">Gasoline</option>
                <option value="Other">Other</option>
            </select>
            <label>Location:</label>
            <input type="text" placeholder="Country" />
            <input type="text" placeholder="City" />
            <label>Phone number:</label>
            <input type='text' placeholder='Phone Number' />
            <button>Submit</button>
        </div>
    );
}
