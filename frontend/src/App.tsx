import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import {AppBar, Toolbar, Container, Button} from '@mui/material';

import HomePage from './pages/HomePage';
import CreateTeamForm from './components/Team/CreateTeamForm';
import GroupListPage from './pages/GroupListPage';
import CreateGroupForm from './components/Group/CreateGroupForm';

const App: React.FC = () => (
  <Router>
    <AppBar position="static">
      <Toolbar>
        <Button component={Link} to="/" color="inherit" sx={{ mr: 2 }}>
          Home
        </Button>
        <Button component={Link} to="/teams/new" color="inherit" sx={{ mr: 2 }}>
          Create Team
        </Button>
        <Button component={Link} to="/groups" color="inherit" sx={{ mr: 2 }}>
          Groups
        </Button>
        <Button component={Link} to="/groups/new" color="inherit">
          Create Group
        </Button>
      </Toolbar>
    </AppBar>

    <Container maxWidth="lg" sx={{ mt: 4 }}>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/teams/new" element={<CreateTeamForm />} />
        <Route path="/groups" element={<GroupListPage />} />
        <Route path="/groups/new" element={<CreateGroupForm />} />
        {/* позже добавим EditGroupForm и GroupDetailsPage */}
      </Routes>
    </Container>
  </Router>
);

export default App;