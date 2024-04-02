import React from 'react';

const SpeedIcon = props => (
  <img
    style={props.style}
    alt="speed"
    src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4KPHN2ZyBmaWxsPSJub25lIiB2aWV3Qm94PSIwIDAgMjAgMjEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiPgo8cmVjdCB3aWR0aD0iMjAiIGhlaWdodD0iMjAuMTc1IiBmaWxsPSJ1cmwoI2EpIi8+CjxkZWZzPgo8cGF0dGVybiBpZD0iYSIgd2lkdGg9IjEiIGhlaWdodD0iMSIgcGF0dGVybkNvbnRlbnRVbml0cz0ib2JqZWN0Qm91bmRpbmdCb3giPgo8dXNlIHRyYW5zZm9ybT0ibWF0cml4KC4wMTU3NjIgMCAwIC4wMTU2MjUgLS4wMDQzODYgMCkiIHhsaW5rOmhyZWY9IiNiIi8+CjwvcGF0dGVybj4KPGltYWdlIGlkPSJiIiB3aWR0aD0iNjQiIGhlaWdodD0iNjQiIHhsaW5rOmhyZWY9ImRhdGE6aW1hZ2UvcG5nO2Jhc2U2NCxpVkJPUncwS0dnb0FBQUFOU1VoRVVnQUFBRUFBQUFCQUVBWUFBQUQ2K2EyZEFBQUFCR2RCVFVFQUFMR1BDL3hoQlFBQUFDQmpTRkpOQUFCNkpnQUFnSVFBQVBvQUFBQ0E2QUFBZFRBQUFPcGdBQUE2bUFBQUYzQ2N1bEU4QUFBQUJtSkxSMFFBQUFBQUFBRDVRN3QvQUFBQUNYQklXWE1BQUFCZ0FBQUFZQUR3YTBMUEFBQUFCM1JKVFVVSDVnZ1BFeWtzRENUbmlRQUFGSEZKUkVGVWVOcnRuWGxBVkZYN3g3L1BuUkVWbVVGRnlpMDFkOFVOeGNRRkVRWVhYRkFaQnEzY2R6T1ZRaVZUVTdMY2MzM1YzRXROeXdIY1FZTkJXZlNseERRMTlDVlN5aDFGWVFBMzRENi9QNWc3RU9NRUdEaldqODgvT25PVyt5eG43ajMzbk9jOEFPV1VVMDQ1NVpSVHp2OUh5TklDbERiMkdsZlcrTmpZUEtzb1c1MnpzV2xUYm91M1pSODBiWXB6c09kUmRuWjBGSWZSd3NaRzdJUVFMS3RhVldvbi9BQnZ6RXBMNDM0WWdNdVptWERFUGZvcU5aVit4dDdjVlltSlZrOXovZVNURXhQdmFhTklHNVNaYVdrOVM0dC8zQUNvUE1kOWhFWlRwNDVjSkloT0hoNlVpa3Y0enMyTlIyRUp0cmk2VWg4c3hhSUdEY3JxK2h6SDIzSDQyalZhVFlQUktqcWFPOUYxNmhJWm1YTlY5S1M3T3Qzanp5TjNhclUzYjFyYVRzWGxsUjBBZGdlNzJubkZLaFRQQXExdVcrMzM4VUZIb1RMdkdUR0NZdmdVWCt2ZUhUZG9EcjByQ0phVzAwaGQvcHkvRVVVZVNCdG9UVlFVa3JnL3Y3bHpwOVd3cDVOemFnY0hwdzQ4bFhxb1cwYUdwY1VzekNzekFLcnM3UG56NEJtdnZTWmNGcitnUWRPbll4d3UwYlFwVThnUjFmQ21yVzFKKytQMStJTlhaV1ZSSWo2bTZjbkozQUV6MGU3MmJmTEZNWjdYb0FFcUlZQTBqUnNiR3p6QlV0WW1KZkUrOUtHRnljbDBGc3R4dmxZdGJvcEZ2S1pCQTVxQ2V2UkJsU29sbHVNcU9zRTVMUTE3TVFIL1c3K2V0MWw1Q0dscjFtU2VEMHZTQnQyN1oybTdXMndBMU5yVVA3NS92TFYxWnZManRWWlA1c3hCWFE2a2kzNStOSnZHSWNMYXVxajJ2SlRYWVA3MTY5U2EvSEgxKysreGpYL0g5em9kMTZTM2NxZWVQWnV4dU52WGJac25KUUdCRkVpaUtMVlRXcmwvN2ROaTdWcFVvcDNjY3VwVVk0ZFBlQVFsckZ1bmZ4WTVNdWp5dEduNVY1clA4MWtRRkEyaWR5ZjRObWxDTldWZjV6YnIwQUV6eEEwVXBGSnhFcDFIdjE2OWFCRTJJYmx1M1NMbFhzeGI0ZkhvRVdwaUxmbXRYR2x6My9yZEo2Y1dMNzQ5OFlqVEVhZEhqMTYySDE3NkFMQ0pWN1ZYdHg4MGlGTFFGVjFYcjZhaFNNQ3QrdlhOR2l5S3YrWEZlajMya1QrQzl1NEZBTm01cjcvT1dLemJxZFgrOTc4bHZYN0pCMER4VUxpNE82bVR1blJCRXh4SC9NaVJtQXFSMnc4ZFNxNDBsR1lybFdiMU80WUFmSnljek8zSW0vejkvREt0SXpLRFRoNDgrTEw4VWVZRFFQcWxaems4bmxQeHh0cTE2SXNjN0JvNzFteURMb2hHNjVRVUxFWUlDNnRYeTl4b3RlekNoZzBQMHlKMDJxRDA5TDhyVDFrTmdNTFlzaXNQNUtwVithQzhxZnpoZSs5aENXL0FXMzUrU0tBbGFHdHZiN2JoT2N4Z3Q2MWI5WTJ5V3NtaXBrMEQ0dDdRYWg4L0xuWEhHQ2l6U1pSTmZLOEo2dmJObTJkZGZCeG85ZWFaTTJZZDc4dXJrZkhzR1c2QUtHYlpNamxaMWErVTJxU0ozbEUzSU9TVHhZdEx5L0V2bTNTS29vT1VscVlmcFBzMXVQcWlSZkpxRlNkWDhtemNHTUgwTzMreFlvVlI3OEk0WWdXZEdEZE9PYTlLaEpoejVveWlnZHN1amFaWnM3S1NzOVFIZ0dLMmFvUkcwN216TUNBM0VUNnhzZGlKUnpTaFpjdkM5ZmdFRnZEOHMyZkZlSElXYkIwZDlVcGRSRkJLUU1DRHNMQ2tiOWJwOVdXbHNLV1E5TkwzakVnS2FUQnpwcVEzaDhBQlEzNzZ5YVRCR3V5RzRPQkFUNFZ0b3VPcFU0cFFqN1hlZlp5ZFMxdXVVaHNBaXIwZUhvUHY5ZXVIdWh3b3VrZEVJQXN5bkxXek02bjRBMytPWnN1WFozUkltMTVqWGVmT21lZDFXVnB0UWtKcEsyWU9Qb0M5WXYyclYwMitUNk1rM3A2VTlMTGtrUFRPOEVpTHRRdDNkamJlR1FvajJmR2FPSUFHNm5SR081Y1NmM3NPWUh2U1BjUDdpSnNiSDZZbWRDY3NERnZnZ0xDS0ZZMFZacUFPYjNqeWhDZnllN1JtL1BpTTF5T2RneGZ0M3YyeURHMkt4a2ZqWTJXbGVQU2doZWcvYVpMMGJZWjEybEc3N1JzM0FtZC8ycndsTzl0UzBpbnV1c2VwUHg0MmpEYlJCcDYrWlF0VzRDYTlWNm1Tc2NKNC9BTFBwMDlwQVAvS05UMDkwM3RFS2tMNm56anhvdGVUdjJqREtxTThWQnFmMXEzRlZyeFl2QndTUWw1d1FGd0J4MHY4d1ROcHVDQ2dFNS9uMDYvQ3JWMGJwQTE2OWl6REdrRFEycldXbHFZd05KQk9pZjN2MzBkZk03OU93dzlNSElZbldIemdnRzJRS2tXOXNVZVBkQi9kYThHVHo1MHI4ZlZLMmtCYXNKSGRFdDJFc2VmT1lSRWNNYmQyN1NJYkdrWXVXNkVYYmZMeHlWaXMyeGwwNXNnUmkxajVGVVF4V3pYQ3AyUC8vdlFNMy9QRW9DQ1RPNms1bHFFT1ltN2VGTCswT2kxMGQzUXM2UUpUQ2VZQWVRc2lncE9ZSTZ6ZHRjdkU4UzBSaDZlNXVSaFB2N0hidkhsOENIN29YbURwMDZDUXBLQ2tzS1VOYjJtS2NyelJqZ2E3R3Uwc01RczM0VktuRHAxOVJ1TG1uVHZ6dnFSaS83Q0xQUUFVOGxpM2l4L01tRUhPbUlVQnZYb1ZMdWRHck9OcUFRSDZMeUtTUTk3LzdETjRpR1A0bXFkbmtRTWh3ZTE3dFZYZnZwWjJSSkg2UDNKcm9mNXZwMDZLTTZvUHZVL3YzYXV3VmFXb040NFo4OEw5R2ZRdTB2RUdPMHAyNVR1SXc1WFpzd3YzUjlWUUQrUDc5TWxia1BMM0w2NGNSWTRVVzF0VmpQZUZoZzNaSDF2dytpKy9tRXhLdnNWN1dMcC92NzZ2VGgzYzJOdmJSRkc1bTcyM1g5ZXVpQkMyMDV0aFllU0YxWWhXS0l3VkRMY3cvU1RkenVDVW9wZFNYejRhalVZamt5bjJQamdqSHI5NmxTYWlFWHJXcXdjbFZGQXowL1djMlRsdk4yeVk5OTZmbkZ6Y1hwVkxWTlhWbDIvZUxId25MZXo0akp3VDkwSlduenBsMG42NSt4M3ZPd2NPWUNHOVMxTUdEalFXTE1Sb3BENStUSHFodFdEdjRKQStMN3l0Vm52dG1qazVpcndEY0hmTVFJTTFhMHdjL3c1ODhlT2RPK0xIVmdIQ1J4TW5tbXR2Vk1ETUhZR0hJaDRyYjkxNmFmNHNJWW9HOTcxeXcvcjFNenBld2dFL0lqczNOL2NzWUwyNTVQRUIzQk1IT2Z2MmJlUG5ZanBlSXZkMTJWMWVNV0VDK3NJR1MrL2VOUmJNd3c3WVZhN01HV0p6OFl2VnE0dVN3K3dBVUw3dkVlaWQycnMzb21CRG81L3pySDdLcXpEZ3d3K0xPK21RRktLUGN2b0xNZTNiaTlmd2hOOGZONDU3V0QwVERwVGVlMjFwUXpYSmlaSW1UREFwbUk3aDZIcjRjS1pURk8yZGVQOStTZnZsc1ZhalpaOTZla3Ayb0dTeFFlNTFSOGVpSEMrUk5TSzg3ZjRWS1NrNEJpZGE4cHhiL3Bmb2l3Kzh2SlNYM2VNRzcralowNngrNWdxVUMxWDkxVzlHUjJNNUhxTzlpNHV4d0FVRGNmbjBhZjFSM2JUZ2hHN2RET3B3bVh2aUpWTnRiNi90QXo5ODQ0M2NOYmxUNVluWHJpRUJ6cWdva3hrcnRFUmYwZC9UVXgrbjg5L2Y1ZGd4eTBwTGxPZXZxQ2dUZjIzbGlUd2tLa3J2RytrYjR0dWpSK0dXSm5jQW0zZ1BsY2FuV3plVGpneUlnWFJRYUJrUWtQZnAzK2Q0aVJ5WDNFVDVyWEhqQ2p1ZU4rRTNoUC94aHo2dWVweDhWWGk0cGVVMFNNVmlNc2V5Lzl5NUprWGphQk45NStwcW5Jc1Z3bVFBVUREWEVRT25UemZwYUF3ZjROR3hzWmxPRVRwdFVHeXNwVlV1TzF6WmxlVnlPb0piaURUZHZLSVVubzMybXpjRFdxMVdXK0Ixek1Ka2JvdE1DM2svT2hvRCtXMTRQT2NSc3AzVzBEWFRYVTdqQUpDMkw2a3l3RFZNbi9taUU2WWdZT1ZLU3l0YTF0ZzhxcUN3ZCt2WFQzcS9OaFowaGdKZU9Ublp2UVYxOW1jN2RsaGFUbk9Ra3Jwai82cFZKdDhuMHpIZTdPVlZyYXFIU3VPVEgyR1Zmd2VvS28rVlhmVDFOWm50VjBFdU9xU21aZzZ6ZXlScmNmU29wUlVzYTRRMmZKbmJtNS8wUFc0ZVVlTlF0MWYzclNWOWZmV1BCZVhodzZpRDNoeFFZSEpxOEd2MmIvd2tkNHFQajFGZjZUOXNqd2IwVllIM1NZbFVIazI3OXV5UjF0QXRyV0JaSVUzNlVBUE5rZGk3dDBtRlVMckUxVFp0c3JTY1JaUG5KLzZFSzFDQ1ZsdTRWQmlGSEhRWU5NajRXWHJtc1pZcjhTUnBWbCtBOW5RL2QxdG9xS1hWS210eTNzelJ5TnpIanpjNzZmdFB0VXV5U1JFUmxwYXoyRnlnYzdUZTFHKzhHdG5VdzhWRld1QVNGS0VWMXRsNU9qbVp4SzRadG5HcmpLN3NtbTBWSFcxcGZjb093NlR2UEUyaE1hWkx1MVFaRlpDOGFkT3JOdWtyQ3V0NHdiL3k3eWRPU0p0d1JuME1VZFpLNTlUSzRBNGQ1RFFhUzNHdlkwY0F6VkV3NkhrM0w2SWo1OC9mVGp4Uzk4allseCt0V2xaSTI5anlFRkZrYnQ2Y2IxQlB2TzdraExwZ0xqanA2dzF2M3A2ZG5mT0Z2Q1kvMjdFRGdOclNzcGVFdStIaGJYZWxaR1VwZGFvRjNxMSsvaGsyaUtHd3Q5NHlWbWlJSGVLa2poMEYzc3UzNE5HOGVlRU9PQkJheUVxK3YveXFJWVZTS2E2NUo2aTduRGtqQzJHSWZPRUNnd1RHdm4xNWpwODF5NlNoaHUxeC9QVHBSekhIMWZ2dDg1ZHMvMm53S0o1S3llZlBteFFFWVRkTmJkWk1nQjUyWk4yMGFlRnk4aE5XOExyTGx5MnR3SXRpM0diVjhTTHlPWG1TMnRKVTFISnlLbllIMGdKS3FIc2Y3ejUrZnBiVzU0V1pnVXlNTUEyNTQvUFVpbjlyMWt5Z2o5QVNiZ1UyT1F5UUQ1OGd1MWYzZGNjY1VzQUtIVUFyY2YyT0hlWUNLM2dLNnVEeGhRc0k1QUNNdlhSSk90cGwwbUUyTmFiMGhRdXRYWG9IRDc1WHE1YWw5U3Nwd2dScXllNm1aeFhKbjVQaFhhK2VISU40R1k0cEZOaEFjLzVrb0FEYWl1b3ltZTN1dk8xZ2N4ZVFkY25waGxuUG5qMElpNktRWXpkdVdGemhNZUoweWhrMUN0YVEwZElhTll6NkdIYmIrREkreFBCKy9USW42M1RCMVdOaUFPeEVhUDdjUU5hR2JjU2w0ZUVJUlNZQ1huK2RodU15YXR2WXlNZm5IcUh2eG81RkRBQjg5cG1sOWF6dTZjcmVmZXJXelQwdGo4VXlLeXR6OVhnUEhhWkZnb0JXZkJZRkl4MWRhVHptSzVXa09LaDZYOTA1STBOUzlJVWxlZzFmb1g1b3FENUo5MGJ3U3N2dDdpa21xbVRxV2lkTzBGNzBRSmNDbXg5S2R1UEJDeGJvYjBUT0RSa1dHR2l1dlZLaldxOStQRzhlamlNRXd6NzkxR2pJbmZnWmJTSWlNZ2JwN2dmUE43KzdWcllRS2F1cUZucm5IRDBLRVNkcGlLZm5pL1lrL1NBRVdweTNmL3kzWlV2QktQemV0NiswcEd3WkF3SFFZQXJXbUI3RkVrY0pMWVU5Rnk4VzFaeHVZeFlQZnM3Y3B3ZGVRNk9TSDFJdExXeHR1M1Y3NTUycVZmK3U0NDE2Zm9TaENMZTJGbmdKdmtYUHYvR2FKOFdvaFVLTzRkdTJTU2RpTEdVb3RPWjNjTjQwUG9FY3hSam9IQjJMYXM3enFSYSs3TlRKcE9BV2RtSE13NGVXVWlzOVBUWjJ6NTZIRDdHWWZtQ3ZIVHRNWWdOTGlPUjNPUjFqZTl6TnpNd0xEY2dQMVJJakNRSzV1QlM1K3hkbitMY2JqdU5WaVBHdFF4L3hvOGhJV0VOR0tMQ2s2MFhlNHBrcFUyeVZLbTl2UkVXbDYzVXVJVzJrbFQwaVphZ3EySHVzcnkrTzhSU0tLSEJtMEFDNWtSZTNOa1F4bjdXY2V2b3BFWmtoSTZVRnF6RmpBT2orcXI3TlJsVVQ5UU1YRnlFQTlUQStmMEdQNGpBU3ZUSXk1RGhKcDNodVhydytEYytmNWNyOE9DNDNzMlpOeTZuNllwQ1RQRWwyYyt0V1hNaE9FcThFQkVCTy9kQzhlbld5d3dJS3JGWXRMNEFoUEZ4WlM5VkR2U2cxRmZWNURyYUxJb1ppQ2RMczdRRTRJQ3kvUCtrWXVqQXB4eUczOWE1ZElOQ3JrMVdoYUdRQlhKZkcxNnpKSVB3cGVPTUMzOFBOOUhTQlQrRkxlbUlhTkNodTVnU0tMTEF5OWc5QkgzZjhobGI3NEFFZlIwMmU3T2NuQlcrYVZKU09YSms1cld0TTdMQUhvMkE5ZEtqRkgyMHZpTGlVanZMODV3VGJodElpN3BHY0xGQkRwTE5OWW1MaGNxcElDYnl5VlN0TEsvQ2laTGhFcG9XOHYydVhlSUIwUWtqMzdwaUp5dmdwSnNac0Eya2RZQkpDc2VyUUlUaUpMWVFmblowellpTGpneHVmUG0xcGZWNFVFdWsxUHRlNnRVbUJPeTBVbXYzdmYzSjA1ZHNVblppSUkvU25XSCsyZ3o5VmJkdlcwZ3I4WGY0OGgrbmV2ZklWai90ZXUyclhyc0Jpa3V4WEJ3ZHhCZnBRMXRPbjdDcXJ6SE92WE1rYUVWNHhaRVZLaXFYbExpM1ltY1BnMktZTkFRdHdxRUNCRjdxTFVZbUpaR3ZyNXFiUmRPakFMQWlpR0I5dnJHRFlESkZuV3EycVBMZEdqWC9yc2UxL0sxS1NyZXpkbFQ2VnQwNU54WEdFMEpnS0ZhUnlTaGVxQzRLam81Q2VYcU1HY1A2ODhaa25ZV2lRM2U1Wmw2ZHgzYnRiV3FGeVNzYlRPcFhxVjVpbVVoVjJQSEw0S0s0OGVKQk9YZmExM0hmaGdpRHRjOU5NMU1PMTV6d2o3L0lxOWg4d3dOSUtsVk15aENld3h5ZlBPYyt4Z2J6eFZuUzBsRHdyUHlTc0prNmpmbkN3U1lNNTFCUExoZ3dCbks5ck5LV3dZbGhPR2VQS0k3bFNKVzZPRHZoVmJSckQ0SWZPSkFzS2tqNGFCNENWODVPajJWMURRb3hwekF3WUkwaENxL3lZcS9UeXNyUjY1ZncxeXVXeXUxbkRCdzJpaHZnQmNmbEw4bExleElvOWMrWlNXbjRXTXVNQU1HYXk3RS9aMkcrYXBveGJjQXU2UEdPR3BSVXM1NjloZC9wVm5HWHFKNXFNcDdoeTRFRGhYTWVtQjBPMjRoUW5QQ2V1M0JCUVlhdFV4WGhmOFBDd3RLTGwvQm5sVjI0WGZlcjI2a1Z1V0VDQkhUb1VMaWVaMkVpMnhkU3ZKZ05BdjFBM091VEdtVE84aW1kQ3B6TlpaeFpYY0FWVVhyNWNTaGhoYWNYTHlmTUR0eGJTeE5CRml3cVg4cS84T1pvZFA1NmVmdUtFVm52V1pCZkRyQU9GTWRRYkQwMERIMmdHemFHUDJyV3pTWTN0ZHZIazZORWxGVmRLdENEdHUwdjVCeXh0eHBlTnBMZGtCNld6K3dpTlQ0R2d6ZUwyMC9HVTM2VnZ4bzQxK2VVYmxzRDV1akJIU0RRZndHSjJBS1NUam9McDVFbitENDlGenI1OWhjdEppU1d3Vzc1Y1N0OWVsS0RHR0wxcHRBL3pvNk9sZ0FzZWptV29jL2l3UmJ4Z1FiZ2ZkNkdVSTBlTWdTZXRNRlBVeDhRVU4zV09kSkJGRE9NcXVMMXNtVW4vcCtDRHI3Nzl0cWpkM0NKdjRSWDI1cjdOVC8zOWVSZGE0RmIrNUVIYVhaTy9TOVBGZzl1M20zc2tHQjEvajJkeXRlQmc3Q00vS1BKRG1IZ1oybEZpTVpJaC9jdmd1YmlCdmdXTzRCbnNJdG5KL0VESXMzUE9vNXphOHUwN2Rwak05ZzJSUGptUHFYcjJKMFZQMm9zY0FGS3NIOFdqQTErWGpvWG5JK1VNVWdUSGZQaExwZnhRcXlJZEwyWEU2TTFoc0I4eHd0SU9lZWxNaEJvSnc0YVo1RkFxWWlBb1pzZU92SGhxNFVMNmdKWkRwVklWN3Bia0hJMzZzMllWOXd4amlYZTJGZVBkbjZxSEJnWFJkOVFYMlFVV0dxUnRWeGxXczgrNmRlakpPdW8vYVpKWnh4Y3pGY3EvblNKektFazVoY05KeFVlKy9CSzU4S09ncVZPaGh3N0IrZG5BZUQ0VzhIeXROc05mNXhMU3h0ZTN1TmN2OFFDUVl2N0V4L0pQNUhIeDhWUVRwN0dpVWFPaTJwVTcvcThwY2lDWXcvQ0hMbVRXRkNEemRYSXFhWEx0RXIvR1NZRVJ0RS8welEzdTNkc2tTVkVocEJVbzloTkdDN0VEQnBRNy92a1k3VEtQQTlHdlR4L2ozMGt3aCtINE45Y1dMOHVHOU85dnNhenEwbmF5NHFUcXNIcU5YcTlVcWxScU5iTmlsNnFKOXgrWm1UYnRldGJRYUZ4ZExXbmNmeUxTSDZCUW5ITy81NTJZbm02MDYzcjNhK3FOV1ZsbGxUMzhoYWs2eEczWTRNTnQyeXJydW4vbXZYdisvTHdUT3YvY2lLSlhoU3AzM09hcWh6bzRLSjFWVmRUOVAvbWs2aUgzT1BYSGJkcFlXcTV5eWltbm5ITEtLZWVmemY4QkY0YjNLRDk2aUpzQUFBQWxkRVZZZEdSaGRHVTZZM0psWVhSbEFESXdNakl0TURndE1UVlVNVGs2TkRFNk5EUXJNREE2TUREakZuMXBBQUFBSlhSRldIUmtZWFJsT20xdlpHbG1lUUF5TURJeUxUQTRMVEUxVkRFNU9qUXhPalEwS3pBd09qQXdra3ZGMVFBQUFDaDBSVmgwWkdGMFpUcDBhVzFsYzNSaGJYQUFNakF5TWkwd09DMHhOVlF4T1RvME1UbzBOQ3N3TURvd01NVmU1QW9BQUFBQVNVVk9SSzVDWUlJPSIvPgo8L2RlZnM+Cjwvc3ZnPgo="
  />
);

export default SpeedIcon;