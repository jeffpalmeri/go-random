const first = () => {
  try {
    const rand = Math.random() * 10;
    if (rand < 5) {
      throw new Error("an error has occured");
    } else {
      return {
        status: 200,
        message: "Success",
      };
    }
  } catch (e) {
    console.log(e);
    return "500";
  }
};

const second = () => {
  try {
    return first();
  } catch (e) {
    console.log("HERE!!!");
    return "OK will this ever happen or no???";
  }
};

const c3 = (n, sum = 1) => {
  for (let k = 2; k <= n; k++) {
    sum *= (n + k) / k;
  }
  return Math.ceil(sum);
};

for (let i = 0; i <= 10; i++) {
  console.log(c3(i));
}

// export function generateDivTags(numberOfTags: number) {
// 	const answer: string[] = [];

// 	update("<div>", numberOfTags - 1, numberOfTags, answer)
// 	return answer
// }

// const update = (str: string, opening: number, closing: number, answer: string[]): void => {
// 	console.log(str);

// 	// if(closing === 1) {
// 	// 	answer.push(str + '</div>');
// 	// }
// 	if(closing === 0) {
// 		answer.push(str)
// 	}

// 	if(opening > 0) {
// 		update(str + '<div>', opening - 1, closing, answer);
// 	}

// 	if(opening < closing) {
// 		update(str + '</div>', opening, closing - 1, answer);
// 	}
// }
