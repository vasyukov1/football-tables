// src/components/Shared/TeamStatRow.tsx
import React from 'react';
import { Team } from '../../types/models';
import { TableRow, TableCell } from '@mui/material';

const TeamStatRow: React.FC<{ team: Team }> = ({ team }) => {
    const goalDifference = team.goalsFor - team.goalsAgainst;

    return (
        <TableRow hover>
            <TableCell component="th" scope="row">
                {team.name}
            </TableCell>
            <TableCell align="right">{team.played}</TableCell>
            <TableCell align="right">{team.wins}</TableCell>
            <TableCell align="right">{team.draws}</TableCell>
            <TableCell align="right">{team.losses}</TableCell>
            <TableCell align="right">{team.goalsFor}</TableCell>
            <TableCell align="right">{team.goalsAgainst}</TableCell>
            <TableCell align="right">{goalDifference}</TableCell>
            <TableCell align="right" sx={{ fontWeight: 'bold' }}>
                {team.points}
            </TableCell>
        </TableRow>
    );
};

export default TeamStatRow;