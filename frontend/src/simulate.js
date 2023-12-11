import React, { useState } from 'react';

const SimulateEvent = () => {
  const [jsonText, setJsonText] = useState('');
  const [simulateId, setsimulateId] = useState('');

  const handleJsonSubmit = () => {
    const apiUrl = `/api/simulate/`;

    fetch(apiUrl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: jsonText,
    })
      .then(response => {
        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }
        return response.json();
      })
      .then(data => {
        console.log('POST request successful:', data);
        // Handle the response data as needed
      })
      .catch(error => {
        console.error('Error making POST request:', error);
      });
  };

  return (
    <div>
      <h2>Workflow</h2>
      <label>
        Simulate ID:
        <input type="text" value={simulateId} onChange={(e) => setsimulateId(e.target.value)} />
      </label>
      <button onClick={handleJsonSubmit}>Submit</button>
    </div>
  );
};

export default SimulateEvent;
