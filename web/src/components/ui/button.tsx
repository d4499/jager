import { ButtonRootProps, Root } from "@kobalte/core/button";
import { styled } from "@macaron-css/solid";
import { Component } from "solid-js";

const KobalteButton: Component<ButtonRootProps> = (props) => (
	<Root {...props} />
);

export const Button = styled(KobalteButton, {
	base: {
		padding: "8px 16px",
		border: "none",
		cursor: "pointer",
	},
	variants: {
		variant: {
			primary: {
				backgroundColor: "white",
				color: "black",
			},
			danger: {
				backgroundColor: "red",
			},
			success: {
				backgroundColor: "green",
			},
		},
	},
});
