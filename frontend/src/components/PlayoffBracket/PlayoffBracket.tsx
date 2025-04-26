import React from 'react';
import { Match } from '../../types/models';
import MatchNode from './MatchNode';
import { Box } from '@mui/material';

const PlayoffBracket: React.FC<{ matches: Match[] }> = ({ matches }) => {
    return (
        <Box display="flex" flexDirection="column" alignItems="center" gap={4}>
            <Box display="flex" gap={4}>
                {/* Полуфиналы */}
                <Box display="flex" flexDirection="column" gap={2}>
                    {matches.filter(m => m.stage === 'semifinal').map(match => (
                        <MatchNode key={match.id} match={match} />
                    ))}
                </Box>

                {/* Финал */}
                <Box display="flex" alignItems="center">
                    {matches.filter(m => m.stage === 'final').map(match => (
                        <MatchNode key={match.id} match={match} />
                    ))}
                </Box>
            </Box>
        </Box>
    );
};

export default PlayoffBracket;