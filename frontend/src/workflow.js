import React, { useState } from 'react';

const Workflow = () => {
  const [jsonText, setJsonText] = useState('');
  const [workflowId, setWorkflowId] = useState('');

  const handleJsonSubmit = () => {
    const apiUrl = `/api/workflow/${workflowId}`;

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
        Workflow ID:
        <input type="text" value={workflowId} onChange={(e) => setWorkflowId(e.target.value)} />
      </label>
      <br />
      <textarea
        rows="10"
        cols="50"
        value={jsonText}
        onChange={(e) => setJsonText(e.target.value)}
        placeholder="Paste JSON text here"
      />
      <br />
      <button onClick={handleJsonSubmit}>Submit</button>
    </div>
  );
};

export default Workflow;
