import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import Workflow from './workflow';
import SimulateEvent  from './simulate';

function App() {
  return (
    <Router>
        <Routes>
          <Route path="/workflow" element={<Workflow />} />
        </Routes>
        <Routes>
          <Route path="/" element={<Workflow />} />
        </Routes>
        <Routes>
          <Route path="/simulate" element={<SimulateEvent />} />
        </Routes>
    </Router>
  );
}

export default App;
