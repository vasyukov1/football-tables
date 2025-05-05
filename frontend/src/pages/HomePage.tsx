import React, { useEffect, useState } from 'react';
import {
    Button,
    Typography,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Paper,
    CircularProgress,
    Box
} from '@mui/material';
import { Link } from 'react-router-dom';
import apiClient from '../api/client';
import { Team } from '../types/models';

const HomePage: React.FC = () => {
    const [teams, setTeams] = useState<Team[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchTeams = async () => {
            try {
                const response = await apiClient.get<Team[]>('/teams');
                setTeams(response.data);
            } catch (err) {
                setError('Failed to load teams');
            } finally {
                setLoading(false);
            }
        };

        fetchTeams();
    }, []);

    const handleRetry = async () => {
        setLoading(true);
        setError('');
        try {
            const response = await apiClient.get<Team[]>('/teams');
            setTeams(response.data);
        } catch (err) {
            setError('Failed to load teams');
        } finally {
            setLoading(false);
        }
    };

    if (loading) {
        return (
            <Box display="flex" justifyContent="center" mt={4}>
                <CircularProgress />
            </Box>
        );
    }

    if (error) {
        return (
            <Box textAlign="center" mt={4}>
                <Typography color="error" gutterBottom>
                    {error}
                </Typography>
                <Button variant="contained" onClick={handleRetry}>
                    Retry
                </Button>
            </Box>
        );
    }

    return (
        <div>
            <Typography variant="h2" gutterBottom>
                Football Tournament
            </Typography>

            <Button
                component={Link}
                to="/teams/new"
                variant="contained"
                size="large"
                sx={{ mb: 4 }}
            >
                Create New Team
            </Button>

            <TableContainer component={Paper}>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>Team Name</TableCell>
                            <TableCell align="right">Played</TableCell>
                            <TableCell align="right">Wins</TableCell>
                            <TableCell align="right">Draws</TableCell>
                            <TableCell align="right">Losses</TableCell>
                            <TableCell align="right">Goals For</TableCell>
                            <TableCell align="right">Goals Against</TableCell>
                            <TableCell align="right">Points</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {teams.map((team) => (
                            <TableRow key={team.id}>
                                <TableCell component="th" scope="row">
                                    {team.name}
                                </TableCell>
                                <TableCell align="right">{team.played}</TableCell>
                                <TableCell align="right">{team.wins}</TableCell>
                                <TableCell align="right">{team.draws}</TableCell>
                                <TableCell align="right">{team.losses}</TableCell>
                                <TableCell align="right">{team.goalsFor}</TableCell>
                                <TableCell align="right">{team.goalsAgainst}</TableCell>
                                <TableCell align="right" sx={{ fontWeight: 'bold' }}>
                                    {team.points}
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </div>
    );
};

export default HomePage;