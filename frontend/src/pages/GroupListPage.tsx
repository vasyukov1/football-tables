// frontend/src/pages/GroupListPage.tsx
import React, { useEffect, useState } from 'react';
import {
  Box, Typography, Button, CircularProgress,
  Table, TableBody, TableCell, TableContainer,
  TableHead, TableRow, Paper
} from '@mui/material';
import { Link } from 'react-router-dom';

import { fetchGroups } from '../api/group';
import { Group } from '../types/models';

const GroupListPage: React.FC = () => {
  const [groups, setGroups] = useState<Group[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    fetchGroups()
    .then(res => {
        console.log(res.data);
        setGroups(res.data);
      })      
      .catch(() => setError('Failed to load groups'))
      .finally(() => setLoading(false));
  }, []);

  if (loading) {
    return (
      <Box textAlign="center" mt={4}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Box textAlign="center" mt={4}>
        <Typography color="error">{error}</Typography>
      </Box>
    );
  }

  if (groups.length === 0) {
    return (
      <Box textAlign="center" mt={4}>
        <Typography>No groups found</Typography>
        <Button component={Link} to="/groups/new" sx={{ mt: 2 }}>
          Create one
        </Button>
      </Box>
    );
  }

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Groups
      </Typography>
      <Button
        component={Link}
        to="/groups/new"
        variant="contained"
        sx={{ mb: 2 }}
      >
        Create New Group
      </Button>

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Group Name</TableCell>
              <TableCell>Teams</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {groups.map(g => (
              <TableRow key={g.id}>
                <TableCell>{g.name}</TableCell>
                <TableCell>
                  {(g.teams ?? []).map((t, idx) => (
                    // Если ты рендеришь несколько элементов, ставь key
                    <span key={t.id ?? idx}>
                      {t.name}{idx < (g.teams?.length||0) - 1 ? ', ' : ''}
                    </span>
                  ))}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
};

export default GroupListPage;
