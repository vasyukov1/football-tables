import React, { useState } from 'react';
import { Button, TextField, Box, Typography, CircularProgress } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import apiClient from '../../api/client';
import { ApiError } from '../../types/models';

const CreateTeamForm: React.FC = () => {
    const [teamName, setTeamName] = useState('');
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        if (!teamName.trim()) {
            setError('Team name is required');
            return;
        }

        try {
            setLoading(true);
            const response = await apiClient.post('/teams', { name: teamName });

            if (response.status === 201) {
                navigate('/');
            }
        } catch (err) {
            const error = err as ApiError;
            // Получаем сообщение из ответа сервера
            setError(error.message || 'Failed to create team');
        } finally {
            setLoading(false);
        }
    };

    return (
        <Box sx={{ maxWidth: 500, mx: 'auto', mt: 4, p: 3 }}>
            <Typography variant="h4" gutterBottom component="div">
                Create New Team
            </Typography>

            <form onSubmit={handleSubmit}>
                <TextField
                    fullWidth
                    label="Team Name"
                    variant="outlined"
                    value={teamName}
                    onChange={(e) => {
                        setTeamName(e.target.value);
                        setError('');
                    }}
                    error={!!error}
                    helperText={error}
                    margin="normal"
                    autoFocus
                />

                <Box sx={{ mt: 3, display: 'flex', gap: 2 }}>
                    <Button
                        type="submit"
                        variant="contained"
                        color="primary"
                        disabled={loading}
                        startIcon={loading && <CircularProgress size={20} />}
                    >
                        {loading ? 'Creating...' : 'Create Team'}
                    </Button>

                    <Button
                        variant="outlined"
                        onClick={() => navigate('/')}
                    >
                        Cancel
                    </Button>
                </Box>
            </form>
        </Box>
    );
};

export default CreateTeamForm;