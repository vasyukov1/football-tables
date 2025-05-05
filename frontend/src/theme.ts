import { createTheme } from '@mui/material/styles';

// const theme = createTheme({
//     palette: {
//         primary: {
//             main: '#2E7D32',
//         },
//         secondary: {
//             main: '#FFC107',
//         },
//     },
//     typography: {
//         fontFamily: [
//             'Roboto',
//             'Arial',
//             'sans-serif',
//         ].join(','),
//     },
// });

const theme = createTheme({
    components: {
        MuiTableRow: {
            styleOverrides: {
                root: {
                    '&:nth-of-type(odd)': {
                        backgroundColor: '#f5f5f5',
                    },
                    '&:hover': {
                        backgroundColor: '#e0e0e0',
                    },
                },
            },
        },
    },
});

export default theme;