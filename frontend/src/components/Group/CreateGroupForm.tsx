// frontend/src/components/Group/CreateGroupForm.tsx
import React, { useEffect, useState } from 'react';
import { Box, Typography, Button, TextField, MenuItem, Select, InputLabel, FormControl, OutlinedInput, Chip, CircularProgress } from '@mui/material';
import { useNavigate } from 'react-router-dom';

import apiClient from '../../api/client';
import { Team } from '../../types/models';
import { createGroup } from '../../api/group';

const CreateGroupForm: React.FC = () => {
  const [name, setName] = useState('');
  const [teams, setTeams] = useState<Team[]>([]);
  const [selectedIds, setSelectedIds] = useState<number[]>([]);
  const [loading, setLoading] = useState(true);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    apiClient.get<Team[]>('/teams')
      .then(res => setTeams(res.data))
      .catch(() => setError('Failed to load teams'))
      .finally(() => setLoading(false));
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setSubmitting(true);
    try {
      await createGroup({ name, teamIds: selectedIds });
      navigate('/groups');
    } catch (err: any) {
      setError(err.message || 'Failed to create group');
    } finally {
      setSubmitting(false);
    }
  };

  if (loading) return <CircularProgress />;

  return (
    <Box maxWidth={500} mx="auto">
      <Typography variant="h4" gutterBottom>Create New Group</Typography>
      <form onSubmit={handleSubmit}>
        <TextField
          fullWidth
          label="Group Name"
          value={name}
          onChange={e => setName(e.target.value)}
          margin="normal"
          required
        />
        <FormControl fullWidth margin="normal">
          <InputLabel id="teams-label">Teams</InputLabel>
          <Select
            labelId="teams-label"
            multiple
            value={selectedIds}
            onChange={e => setSelectedIds(e.target.value as number[])}
            input={<OutlinedInput label="Teams" />}
            renderValue={(selected) => (
              <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                {selected.map(id => {
                  const t = teams.find(t => t.id === id);
                  return <Chip key={id} label={t?.name} />;
                })}
              </Box>
            )}
          >
            {teams.map(t => (
              <MenuItem key={t.id} value={t.id}>
                {t.name}
              </MenuItem>
            ))}
          </Select>
        </FormControl>

        {error && <Typography color="error" sx={{ mb: 2 }}>{error}</Typography>}

        <Box display="flex" gap={2} mt={2}>
          <Button
            type="submit"
            variant="contained"
            disabled={submitting}
            startIcon={submitting ? <CircularProgress size={20} /> : null}
          >
            {submitting ? 'Creating...' : 'Create Group'}
          </Button>
          <Button variant="outlined" onClick={() => navigate('/groups')}>Cancel</Button>
        </Box>
      </form>
    </Box>
  );
};

export default CreateGroupForm;
