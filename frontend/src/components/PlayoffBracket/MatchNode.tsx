import React from 'react';
import { Match } from '../../types/models';
import { Box, Typography, Paper } from '@mui/material';

const MatchNode: React.FC<{ match: Match }> = ({ match }) => {
    return (
        <Paper elevation={3} sx={{ p: 2, mb: 2, minWidth: 200 }}>
            <Box textAlign="center">
                <Typography variant="subtitle2">{match.stage.toUpperCase()}</Typography>
                <Box display="flex" flexDirection="column" gap={1} mt={1}>
                    <Box display="flex" justifyContent="space-between">
                        <span>{match.team1.name}</span>
                        <strong>{match.score1}</strong>
                    </Box>
                    <Box display="flex" justifyContent="space-between">
                        <span>{match.team2.name}</span>
                        <strong>{match.score2}</strong>
                    </Box>
                </Box>
            </Box>
        </Paper>
    );
};

export default MatchNode;