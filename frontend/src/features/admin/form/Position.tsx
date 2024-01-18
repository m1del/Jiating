import React from 'react';

type PositionProps = {
    value: string;
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

const Position: React.FC<PositionProps> = ({ value, onChange }) => {
    return (
        <div>
            <label>
                Position
            </label>
            <input
                type="text"
                id="position"
                value={value}
                onChange={onChange}
                className='mt-1 block w-full rounded-md border-gray-100 shadow-sm focus:border-indigo-300 
                focus:ring focus:ring-indigo-200 focus:ring-opacity-50'
                placeholder='Enter Position'
            />
        </div>
    )
};

export default Position;