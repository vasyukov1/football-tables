import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import {AppBar, Toolbar, Container, Button} from '@mui/material';
import HomePage from './pages/HomePage';
import CreateTeamForm from './components/Team/CreateTeamForm';

const App: React.FC = () => {
  return (
      <Router>
        <AppBar position="static">
          <Toolbar>
            <Button
                component={Link}
                to="/"
                color="inherit"
                sx={{ mr: 2 }}
            >
              Home
            </Button>
            <Button
                component={Link}
                to="/teams/new"
                color="inherit"
            >
              Create Team
            </Button>
          </Toolbar>
        </AppBar>

        <Container maxWidth="lg" sx={{ mt: 4 }}>
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/teams/new" element={<CreateTeamForm />} />
          </Routes>
        </Container>
      </Router>
  );
};

export default App;